package api_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/ipfs/go-log"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/thealonemusk/WarpNet/api"
	client "github.com/thealonemusk/WarpNet/api/client"
	"github.com/thealonemusk/WarpNet/pkg/blockchain"
	"github.com/thealonemusk/WarpNet/pkg/logger"
	"github.com/thealonemusk/WarpNet/pkg/node"
)

var _ = Describe("API", func() {

	Context("Binds on socket", func() {
		It("sets data to the API", func() {
			d, _ := ioutil.TempDir("", "xxx")
			defer os.RemoveAll(d)
			os.MkdirAll(d, os.ModePerm)
			socket := filepath.Join(d, "socket")

			c := client.NewClient(client.WithHost("unix://" + socket))

			token := node.GenerateNewConnectionData().Base64()
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			l := node.Logger(logger.New(log.LevelFatal))

			e, _ := node.New(node.FromBase64(true, true, token, nil, nil), node.WithStore(&blockchain.MemoryStore{}), l)
			e.Start(ctx)

			e2, _ := node.New(node.FromBase64(true, true, token, nil, nil), node.WithStore(&blockchain.MemoryStore{}), l)
			e2.Start(ctx)

			go func() {
				err := API(ctx, fmt.Sprintf("unix://%s", socket), 10*time.Second, 20*time.Second, e, nil, false)
				Expect(err).ToNot(HaveOccurred())
			}()

			Eventually(func() error {
				return c.Put("b", "f", "bar")
			}, 10*time.Second, 1*time.Second).ShouldNot(HaveOccurred())

			Eventually(c.GetBuckets, 100*time.Second, 1*time.Second).Should(ContainElement("b"))

			Eventually(func() string {
				d, err := c.GetBucketKey("b", "f")
				if err != nil {
					fmt.Println(err)
				}
				var s string

				d.Unmarshal(&s)
				return s
			}, 10*time.Second, 1*time.Second).Should(Equal("bar"))
		})
	})
})
