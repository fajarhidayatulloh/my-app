package infrastructures

import (
	"fmt"
	"os"
	"strconv"

	consulapi "github.com/hashicorp/consul/api"
	log "github.com/sirupsen/logrus"
)

//IServiceDiscovery provides an interface for getting data out of Consul
type IServiceDiscovery interface {
	// Get a Service from consul
	Service(string, string) ([]*consulapi.ServiceEntry, error)
	// Register a service with local agent
	Register(string, string, string) error
	// Deregister a service with local agent
	DeRegister(string) error
}

// ServiceDiscovery struct
type ServiceDiscovery struct {
	Consul *consulapi.Client
}

//Init returns a Client interface for given consul address
func serviceDiscoveryConn() (*ServiceDiscovery, error) {
	config := consulapi.DefaultConfig()
	discovery, err := consulapi.NewClient(config)
	if err != nil {
		log.WithFields(log.Fields{
			"action": "open consul connection",
			"event":  "consul_conn_open_err",
		}).Error(err)
		os.Exit(0)
	}
	return &ServiceDiscovery{Consul: discovery}, nil
}

// Register a service with consul local agent
func (c *ServiceDiscovery) Register(serviceName, serviceHostName, servicePort string) error {
	consul, _ := serviceDiscoveryConn()

	var reg = new(consulapi.AgentServiceRegistration)
	// reg.ID = uuid.New().String()
	reg.ID = serviceName
	reg.Name = serviceName
	reg.Address = serviceHostName
	port, _ := strconv.Atoi(servicePort[1:len(servicePort)])
	reg.Port = port

	reg.Check = new(consulapi.AgentServiceCheck)
	reg.Check.HTTP = fmt.Sprintf("http://%s:%v/v1/ping", serviceHostName, port)
	reg.Check.Interval = "5s"
	reg.Check.Timeout = "3s"

	return consul.Consul.Agent().ServiceRegister(reg)
}

// DeRegister a service with consul local agent
func (c *ServiceDiscovery) DeRegister(id string) error {
	consul, _ := serviceDiscoveryConn()

	return consul.Consul.Agent().ServiceDeregister(id)
}

// Service return a service
func (c *ServiceDiscovery) Service(service, tag string) ([]*consulapi.ServiceEntry, error) {
	consul, _ := serviceDiscoveryConn()

	passingOnly := true
	addrs, _, err := consul.Consul.Health().Service(service, tag, passingOnly, nil)
	if len(addrs) == 0 && err == nil {
		return nil, fmt.Errorf("service (%s) was not found", service)
	}
	if err != nil {
		return nil, err
	}
	return addrs, nil
}
