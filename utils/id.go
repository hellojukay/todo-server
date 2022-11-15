package utils

import (
	"log"

	"github.com/bwmarrin/snowflake"
)

var (
	node *snowflake.Node
)

func init() {
	tmp, err := snowflake.NewNode(1)
	if err != nil {
		log.Fatalln(err)
	}
	node = tmp
}
func NextID() int64 {
	return node.Generate().Int64()
}
