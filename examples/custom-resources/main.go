/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Note: the example only works with the code within the same release/branch.
package main

import (
	"context"
	"flag"
	"fmt"
	//apiv1 "k8s.io/api/core/v1"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)

func main() {
	var cluster1/*, cluster2 */ *string
	var kubeconfig1/*, kubeconfig2 */ *string
    var cr *string

	cluster1 = flag.String("cluster1", "", "raw cluster name ")
	// cluster2 = flag.String("cluster2", "" , "to be compared cluster")

    cr = flag.String("cr", "", "cr type")
	flag.Parse()

	file1 := fmt.Sprintf("cluster.%s", *cluster1)
	// file2 := fmt.Sprintf("cluster.%s.yaml", *cluster2)

	if home := homedir.HomeDir(); home != "" {
		kubeconfig1 = flag.String("kubeconfig1", filepath.Join(home, ".kube", file1), "(optional) absolute path to the kubeconfig file")
		// kubeconfig2 = flag.String("kubeconfig2", filepath.Join(home, ".kube", file2), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig1 = flag.String("kubeconfig1", "", "absolute path to the kubeconfig1 file")
		// kubeconfig2 = flag.String("kubeconfig2", "", "absolute path to the kubeconfig2 file")
	}


	config1, err := clientcmd.BuildConfigFromFlags("", *kubeconfig1)
	if err != nil {
		panic(err)
	}

	// config2, err := clientcmd.BuildConfigFromFlags("", *kubeconfig2)
	// if err != nil {
	// 	panic(err)
	// }

	client1, err := dynamic.NewForConfig(config1)
	if err != nil {
		panic(err)
	}

	// client2, err := dynamic.NewForConfig(config2)
	// if err != nil {
	// 	panic(err)
	// }

	customResource := schema.GroupVersionResource{Group: "cloudmesh.alipay.com", Version: "v1", Resource: *cr}

	list, err := client1.Resource(customResource).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	// ch := make(chan string, 200)
	for _, d := range list.Items {
		//replicas, found, err := unstructured.NestedInt64(d.Object, "spec", "replicas")
		//if err != nil || !found {
		//	fmt.Printf("Replicas not found for deployment %s: error=%s", d.GetName(), err)
		//	continue
		//}
		fmt.Printf("%s/%s\n", d.GetNamespace(), d.GetName())
		// go match(client2, d.GetNamespace(), d.GetName(), *cluster2, ch)
	}

    // for i := 0; i < len(list.Items); i++ {
		// fmt.Println(<-ch)
    // }
}

// func match(client dynamic.Interface, namespace, name, cluster string, ch chan string)  {

// 	customResource := schema.GroupVersionResource{Group: "cloudmesh.alipay.com", Version: "v1", Resource: *cr}
// 	_, err := client.Resource(customResource).Namespace(namespace).Get(context.TODO(), name, metav1.GetOptions{})
// 	if err != nil {
// 		ch <- fmt.Sprintf("%v/%v not found in %v", namespace, name, cluster)
// 	} else {
// 		ch <- fmt.Sprintf("%v/%v found", namespace, name)
// 	}
// }
