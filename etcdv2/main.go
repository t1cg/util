package etcdv2

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"time"
	//t1cg library
	"github.com/t1cg/util/apperror"
	//coreos etcd library
	"github.com/coreos/etcd/clientv3"
)

const (
	etcdStart string = "etcd"
	etcdStop  string = "pkill"
)

//startService starts the ETCD service
func startService() {
	c := exec.Command(etcdStart)

	c.Start()
}

//stopService stops the ETCD service after testing is complete
func stopService() {
	c := exec.Command(etcdStop, etcdStart)

	c.Start()
}

//PutTheValue takes in key value pairs and assigns them through the ETCD
//service.
func PutTheValue(key []string, value []string, aeCh chan *apperror.AppInfo) {

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
		a := &apperror.AppInfo{Msg: err}
		a.LogError(a.Error())
		aeCh <- a
		return
	}
	defer cli.Close()

	for i := 0; i <= len(key); i++ {

		_, err = cli.Put(context.TODO(), key[i], value[i])
		if err != nil {
			log.Fatal(err)
			a := &apperror.AppInfo{Msg: err}
			a.LogError(a.Error())
			aeCh <- a
			return
		}
	}
}

//GetTheValue returns the value for a chosen key that is entered as a parameter.
func GetTheValue(key string, valueCh chan map[string]string, aeCh chan *apperror.AppInfo) {

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
		a := &apperror.AppInfo{Msg: err}
		a.LogError(a.Error())
		aeCh <- a
		return
	}
	defer cli.Close()

	requestTimeout := 5 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		log.Fatal(err)
		a := &apperror.AppInfo{Msg: err}
		a.LogError(a.Error())
		aeCh <- a
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)

		var finalKv = make(map[string]string)
		finalKv[string(ev.Key)] = string(ev.Value)

		valueCh <- finalKv
	}
}

//GetThePrefix returns all Key-Value pairs that begin with a common prefix.
func GetThePrefix(prefix string, valueCh chan map[string]string, aeCh chan *apperror.AppInfo) {

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
		a := &apperror.AppInfo{Msg: err}
		a.LogError(a.Error())
		aeCh <- a
		return
	}
	defer cli.Close()

	requestTimeout := 5 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)

	resp, err := cli.Get(ctx, prefix, clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend))
	cancel()
	if err != nil {
		log.Fatal(err)
		a := &apperror.AppInfo{Msg: err}
		a.LogError(a.Error())
		aeCh <- a
		return
	}
	var finalKv = make(map[string]string)
	for _, ev := range resp.Kvs {
		finalKv[string(ev.Key)] = string(ev.Value)
	}
	valueCh <- finalKv
}
