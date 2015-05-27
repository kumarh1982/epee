package epee

import (
	"fmt"
)

func findRegisteredBrokers(zk ZookeeperClient) ([]string, error) {
	paths, err := zk.List("/brokers/ids")

	if err != nil {
		return []string{}, err
	}

	fullPaths := make([]string, 0)

	for _, p := range paths {
		data := make(map[string]interface{})
		err := zk.Get(p, data)

		if err != nil {
			return []string{}, err
		}

		fullPaths = append(fullPaths, fmt.Sprintf("%s:%0.0f", data["host"], data["port"]))
	}

	return fullPaths, nil
}
