package main

import (
	"io/ioutil"
	"log"

	"github.com/domluna/dogo"
)

type CoreOSClient struct {
	Client *dogo.Client
}

func (c *CoreOSClient) createCoreOS(name string, opts *dogo.CreateDropletOpts) error {
	opts.Name = name
	_, err := c.CreateDroplet(opts)
	if err != nil {
		return err
	}
	log.Printf("Succesfully create droplet name: %s\n", name)
	return nil
}

func (c *CoreOSClient) createCoreDroplets(opts *dogo.CreateDropletOpts) (error) {
	err := c.createCoreOS("coreos-4", opts)
	if err != nil {
		return err
	}
	err = c.createCoreOS("coreos-5", opts)
	if err != nil {
		return err
	}
	err = c.createCoreOS("coreos-6", opts)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	client, err := dogo.NewClient("")
	if err != nil {
		log.Printf("Failure creating client\n")
		return
	}
	coreClient := &CoreOSClient{
		Client: client,
	}

	file, err := ioutil.ReadFile("cloud-config.yaml")
	if err != nil {
		log.Printf("Failed opening the cloud-config file")
		return
	}
	opts := &dogo.CreateDropletOpts{
		Region:            "nyc3",
		Size:              "512mb",
		Image:             5914637,
		Keys:           []string{"136188"},
		PrivateNetworking: true,
		UserData:          string(file),
	}

	coreClient.createCoreDroplets(opts)
	log.Printf("All good\n")
}
