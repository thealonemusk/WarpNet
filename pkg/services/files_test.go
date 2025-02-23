package services_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/ipfs/go-log"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/thealonemusk/WarpNet/pkg/blockchain"
	"github.com/thealonemusk/WarpNet/pkg/logger"
	node "github.com/thealonemusk/WarpNet/pkg/node"
	. "github.com/thealonemusk/WarpNet/pkg/services"
)

var _ = Describe("File services", func() {
	token := node.GenerateNewConnectionData(25).Base64()

	logg := logger.New(log.LevelError)
	l := node.Logger(logg)

	e2, _ := node.New(
		node.WithDiscoveryInterval(10*time.Second),
		node.WithNetworkService(AliveNetworkService(2*time.Second, 4*time.Second, 15*time.Minute)),
		node.FromBase64(true, true, token, nil, nil), node.WithStore(&blockchain.MemoryStore{}), l)

	Context("File sharing", func() {
		It("sends and receive files between two nodes", func() {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			fileUUID := "test"

			f, err := ioutil.TempFile("", "test")
			Expect(err).ToNot(HaveOccurred())

			defer os.RemoveAll(f.Name())

			ioutil.WriteFile(f.Name(), []byte("testfile"), os.ModePerm)

			// First node expose a file
			opts, err := ShareFile(logg, 10*time.Second, fileUUID, f.Name())
			Expect(err).ToNot(HaveOccurred())

			opts = append(opts, node.FromBase64(true, true, token, nil, nil), node.WithStore(&blockchain.MemoryStore{}), l)
			e, _ := node.New(opts...)

			e.Start(ctx)
			e2.Start(ctx)

			Eventually(func() string {
				ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
				defer cancel()

				f, err := ioutil.TempFile("", "test")
				Expect(err).ToNot(HaveOccurred())

				defer os.RemoveAll(f.Name())

				ll, _ := e2.Ledger()
				ll1, _ := e.Ledger()
				By(fmt.Sprint(ll.CurrentData(), ll.LastBlock().Index, ll1.CurrentData()))
				ReceiveFile(ctx, ll, e2, logg, 2*time.Second, fileUUID, f.Name())
				b, _ := ioutil.ReadFile(f.Name())
				return string(b)
			}, 190*time.Second, 1*time.Second).Should(Equal("testfile"))
		})
	})
})
