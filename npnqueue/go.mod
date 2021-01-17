module github.com/kyleu/libnpn/npnqueue

go 1.14

require (
	github.com/kyleu/libnpn/npncore v0.0.0 // npn
	github.com/Shopify/sarama v1.27.1
)

replace github.com/kyleu/libnpn/npncore => ../npncore
