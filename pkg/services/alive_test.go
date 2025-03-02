package services_test

import (
	"context"
	"time"

	"github.com/ipfs/go-log"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/thealonemusk/WarpNet/pkg/blockchain"
	"github.com/thealonemusk/WarpNet/pkg/logger"
	node "github.com/thealonemusk/WarpNet/pkg/node"
	. "github.com/thealonemusk/WarpNet/pkg/services"
)

var _ = Describe("Alive service", func() {
	token := node.GenerateNewConnectionData().Base64()

	logg := logger.New(log.LevelError)
	l := node.Logger(logg)

	opts := append(
		Alive(5*time.Second, 100*time.Second, 15*time.Minute),
		node.WithDiscoveryInterval(10*time.Second),
		node.FromBase64(true, true, token, nil, nil),
		l)

	Context("Aliveness check", func() {
		It("detect both nodes alive after a while", func() {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			e2, _ := node.New(append(opts, node.WithStore(&blockchain.MemoryStore{}))...)
			e1, _ := node.New(append(opts, node.WithStore(&blockchain.MemoryStore{}))...)

			e1.Start(ctx)
			e2.Start(ctx)

			ll, _ := e1.Ledger()

			ll.Persist(ctx, 5*time.Second, 100*time.Second, "t", "t", "test")

			matches := And(ContainElement(e2.Host().ID().String()),
				ContainElement(e1.Host().ID().String()))

			index := ll.LastBlock().Index
			Eventually(func() []string {
				ll, err := e1.Ledger()
				if err != nil {
					return []string{}
				}
				return AvailableNodes(ll, 15*time.Minute)
			}, 100*time.Second, 1*time.Second).Should(matches)

			Expect(ll.LastBlock().Index).ToNot(Equal(index))
		})
	})

	Context("Aliveness Scrub", func() {
		BeforeEach(func() {
			opts = append(
				Alive(10*time.Second, 30*time.Second, 15*time.Minute),
				node.WithDiscoveryInterval(10*time.Second),
				node.FromBase64(true, true, token, nil, nil),
				l)
		})

		It("cleans up after a while", func() {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			e2, _ := node.New(append(opts, node.WithStore(&blockchain.MemoryStore{}))...)
			e1, _ := node.New(append(opts, node.WithStore(&blockchain.MemoryStore{}))...)

			e1.Start(ctx)
			time.Sleep(5 * time.Second)
			e2.Start(ctx)

			ll, _ := e1.Ledger()

			ll.Persist(ctx, 5*time.Second, 100*time.Second, "t", "t", "test")

			matches := And(ContainElement(e2.Host().ID().String()),
				ContainElement(e1.Host().ID().String()))

			index := ll.LastBlock().Index
			Eventually(func() []string {
				ll, err := e1.Ledger()
				if err != nil {
					return []string{}
				}
				return AvailableNodes(ll, 15*time.Minute)
			}, 120*time.Second, 1*time.Second).Should(matches)

			Expect(ll.LastBlock().Index).ToNot(Equal(index))
			index = ll.LastBlock().Index

			Eventually(func() []string {
				ll, err := e1.Ledger()
				if err != nil {
					return []string{}
				}
				return AvailableNodes(ll, 15*time.Minute)
			}, 360*time.Second, 1*time.Second).Should(BeEmpty())

			Expect(ll.LastBlock().Index).ToNot(Equal(index))
			index = ll.LastBlock().Index

			Eventually(func() []string {
				ll, err := e1.Ledger()
				if err != nil {
					return []string{}
				}
				return AvailableNodes(ll, 15*time.Minute)
			}, 60*time.Second, 1*time.Second).Should(matches)
			Expect(ll.LastBlock().Index).ToNot(Equal(index))

		})
	})
})
