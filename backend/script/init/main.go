package main

import (
	"flag"
	"github.com/lwl1989/kube-ssh/backend/internal/config"
	"github.com/lwl1989/kube-ssh/backend/internal/model"
	"github.com/lwl1989/kube-ssh/backend/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var configFile = flag.String("f", "../../etc/api.yaml", "the config file")

func main() {
	var c config.Config
	var err error
	conf.MustLoad(*configFile, &c)
	_ = logx.SetUp(c.Log)
	ctx := svc.NewServiceContext(c)
	//err = ctx.DefaultDb.Write().Exec("CREATE DATABASE k8s_manager DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci").Error
	//if err != nil {
	//	logx.Infof("init database error: %+v", err)
	//}
	err = ctx.DefaultDb.Write().Migrator().CreateTable(model.Cluster{})
	if err != nil {
		logx.Infof("init database error: %+v", err)
	}
	err = ctx.DefaultDb.Write().Migrator().CreateTable(model.UserManager{})
	if err != nil {
		logx.Infof("init database error: %+v", err)
	}
	err = ctx.DefaultDb.Write().Migrator().CreateTable(model.UserWhite{})
	if err != nil {
		logx.Infof("init database error: %+v", err)
	}
	err = ctx.DefaultDb.Write().Exec("INSERT INTO `k8s_cluster` (`id`, `name`, `namespaces`, `config`, `server_api`, `status`, `ctime`, `mtime`, `operator_id`) VALUES (1, 'local-本地调试集群', 'default', 'apiVersion: v1\\nclusters:\\n- cluster:\\n    #certificate-authority: /Users/yt-it/.minikube/ca.crt\\n    insecure-skip-tls-verify: true\\n    extensions:\\n    - extension:\\n        last-update: Tue, 02 Jan 2024 15:52:21 CST\\n        provider: minikube.sigs.k8s.io\\n        version: v1.32.0\\n      name: cluster_info\\n    server: http://192.168.40.7:9002\\n  name: minikube\\ncontexts:\\n- context:\\n    cluster: minikube\\n    extensions:\\n    - extension:\\n        last-update: Tue, 02 Jan 2024 15:52:21 CST\\n        provider: minikube.sigs.k8s.io\\n        version: v1.32.0\\n      name: context_info\\n    namespace: default\\n    user: minikube\\n  name: minikube\\ncurrent-context: minikube\\nkind: Config\\npreferences: {}\\nusers:\\n- name: minikube\\n  user:\\n    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURJVENDQWdtZ0F3SUJBZ0lCQWpBTkJna3Foa2lHOXcwQkFRc0ZBREFWTVJNd0VRWURWUVFERXdwdGFXNXAKYTNWaVpVTkJNQjRYRFRJME1ERXdNVEEzTlRJeE1Wb1hEVEkzTURFd01UQTNOVEl4TVZvd01URVhNQlVHQTFVRQpDaE1PYzNsemRHVnRPbTFoYzNSbGNuTXhGakFVQmdOVkJBTVREVzFwYm1scmRXSmxMWFZ6WlhJd2dnRWlNQTBHCkNTcUdTSWIzRFFFQkFRVUFBNElCRHdBd2dnRUtBb0lCQVFERTRTcGNVa2xXMUUvREZXcm80VUJxcWVKMGFNWm4KVmFPaS9la3NXSmgyZlRxV0xHbzEvYWxhTEl1NGNTTTBlLytxMi9pbWt1WGhvRzU5MUpvbkJkYnN5aE1IbDBEdwovTTBqOTUwMHpxMmtXN1hLYjJ5VU9uUktDTUlvbHNoRVlDY2dPYUpWNnkzQ0dEdC9tTklnMEdwajZTRnhWZ0ZICktXSDN6ZDh1dlVHOTc3QnNoTFdNaUJBR1VZaWR6dGZzUkZkVEsvajZJaXRWVCtlQlRsRGJJNXdpOTJKS0kxYjgKY0ROT1FzK2d2c3oyVGw4dzdaYnVvd2s3RzY1SVJndzhLNEZoNlhwUGRkaVpmZzR4OEJJSVZTWWdRQnlxeGJKdQpvb3dLY1lnVG1ZVm9wcndJWlY5azdHU09YNXRBMjlHS0ttWStDejZkR1FYUWNkeXpwUDJGRUJzckFnTUJBQUdqCllEQmVNQTRHQTFVZER3RUIvd1FFQXdJRm9EQWRCZ05WSFNVRUZqQVVCZ2dyQmdFRkJRY0RBUVlJS3dZQkJRVUgKQXdJd0RBWURWUjBUQVFIL0JBSXdBREFmQmdOVkhTTUVHREFXZ0JURzNvbXRVNkJwNUZOSlBjTUw2YnUzNVZ5TQpwekFOQmdrcWhraUc5dzBCQVFzRkFBT0NBUUVBZ1FtWUJrZTZJTlNrYjRxV29jM1JxODlWK0VTclZQVE5JRXFVCnJFWi80YnF1WjcvRWdVWVZ2MEdTdUt6V2NuWEZLTDhISXZIUXRsVmM0YStqanJMcy9VV1pRMC9IejNWR2VBS0UKcC90U0FpWWRiSW93clk0b3EvaVhnQTdCeHhsR09VS08rTTVWTmNwKzdSdTEyZkFwVGRxdy9SaklLK1d2anhDQwpNZGcwVU1ZMXdXcUp0R29yZVZ2UVZTQXk3Wm13RlI0VHF4b1BJVFNHV2lSMzZ3Vkx1VTV6dlFJbDZXTFRhVVNiCnFxY1lMVDB4ZU53MFdNemRpUkNkYTRrSlJvaEVldDRJUzdXZ3g5Si9VdkNNaTdLaXNNM3A5MFB1b3o1MHRXWFQKaEVDYjhkcnc0Nzk0ZitFVlMvQTBOVCtNbmVJK2lyOFZMekxzQkx2UDR0WlNCQUQ4Tmc9PQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==\\n    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb3dJQkFBS0NBUUVBeE9FcVhGSkpWdFJQd3hWcTZPRkFhcW5pZEdqR1oxV2pvdjNwTEZpWWRuMDZsaXhxCk5mMnBXaXlMdUhFak5Idi9xdHY0cHBMbDRhQnVmZFNhSndYVzdNb1RCNWRBOFB6TkkvZWROTTZ0cEZ1MXltOXMKbERwMFNnakNLSmJJUkdBbklEbWlWZXN0d2hnN2Y1alNJTkJxWStraGNWWUJSeWxoOTgzZkxyMUJ2ZSt3YklTMQpqSWdRQmxHSW5jN1g3RVJYVXl2NCtpSXJWVS9uZ1U1UTJ5T2NJdmRpU2lOVy9IQXpUa0xQb0w3TTlrNWZNTzJXCjdxTUpPeHV1U0VZTVBDdUJZZWw2VDNYWW1YNE9NZkFTQ0ZVbUlFQWNxc1d5YnFLTUNuR0lFNW1GYUthOENHVmYKWk94a2psK2JRTnZSaWlwbVBncytuUmtGMEhIY3M2VDloUkFiS3dJREFRQUJBb0lCQUMzdmgwMnhHVkY4Q2Z3dgpkQmxQN1JLMS9wTkFtd0lqTmlIaWNsUVplOEV4cU1pL0ppemd1WEhEc1BuZzArRDhDWVFZL09RSXBFQkhpV0FzCmlhY1BNcjFlekovWng1b3lzYjV4bUtsb2k1VXNuTGJWMXBTakt0elhQRTN5R0ZuendVMUFoVUxjczNsMDQvVDYKZUJTVjdDelJpUzhEYlJyb2FlWkNqNDg5TXlpWWFoejVyZ0duYmd2OHFXWW9yWWFGOWpoK2ZaT0Q2Zm1PWnA5bApMZzBuWFZzZDJnOGhQbFd6cEIycThWdklJdE5kMXp1cG54dldDclNkNnEwOXFlbnlxbG9GU01SU0RqRGQ1Mld1CmNscGVnTng2R0J5RzRseTlhdldWTk5pRW9VQWlRbXBzQVlwRkg1d1NKclY4R2xMYXh1Zk5EYk5od29PRHNvOUEKS0RnSmtERUNnWUVBOFZvelRqVnoydHNUSWVYOWZVZDNuYjRhUGZuS2oyMGZ5MUplbmhRN2lNSTBlblVwVjltbAorUGxlNUlwczIxWTJZNVFYS2l0YWdrV0hQZStmL25IMUNSTmJ4VkIyYi9sV2lyWHBzeUVQVVJOdk5nYklmRXFDClRYdXFyVFozV1QwSlgyRUFhdVpnand6UC9KYVRiYXlYWXRCalJuUkRmSEt6OGdxUWV2dVZ1b2tDZ1lFQTBOUUMKR1FRZG10b2lYbThPYXI3K0dkdHdzRy9tNExWTmx1dmpiNUhMV3RkS1UrSmh1TFFvK1NJeDVzb0U0emdNYThJTwpwZHlZTVZTUUZMKzB0NWV0Rk1QUjJabTJXdmN0VVpFd2NVVzNVbEFuRDcrTDJxY3ZYZ2hOb3ZlZ3ZyZEltdmowCkppRWRqcXNBUzlsbHhublRFbUxLbFVxY2pJcmtJeGtqWXJVY2F4TUNnWUVBdkdKZGZZUTNZL0p2b3B2MEdsODQKUElYdjBjUXhtWFhoeFVBTDNuT0liSnk1ZllRSnV5cUZaQ3F4S00zclhlQ1RIM0t1Q2hwQTBVSVg1LzRyOGQxZApGN0ptaFVMaXowL2Rmdk95OEVDenhlTFhnV0lXQnYzWmEwVkYyV0dVRXJHVHFVRDdwSFVobFVhNDZUMVc0ZG8vClo3K2tYWS9PUlVyNnJjZ1ZNZ2xCdTVFQ2dZQmw1ZjI4RFRrUTliM3RqSTFoWXg4RXFRSmM1YzJuK25BSTQ4UFEKRGpsSGMyUXVlSG1zc2lTSUpMcHEza3J5UU1nMjBMTnJGYkFoNmh5QU0yZFFhcStuUVVJbHh3Nm5acE56aU1BMQpsWW8xblN6aVQxcEQ4RzU1bU4yaFZ1bldCZ05rczNRWEl2T1VTVGJVekJrUWR5T2FoaUJLSnVVcTR4OGRUVWZxCkEyd0Jod0tCZ0JNUlQ2L2hyVExhcXpNRDIvTm41OERRQWN3RmdYWGlIa24yRVRIRW5mUTJHUjgrdzNBN3doZFgKTC9JdEt6YVVSY2haWXRadW9LVlZkQkk2T1JCZUZGbXo2Q20xZWQxRGFUSEJZUUhabHNJY3ppRFdjWG16WlpqUwpheEEwZCs2eGdrNXo4UnJKYUZEak5MUWc5dzJONGx3c2w4dkhyL3d0bTJXQnR2RnhlUTdNCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==', 'http://192.168.40.7:9002', 1, 1704444520, 1704444520, 1);").Error
	if err != nil {
		logx.Infof("init database error: %+v", err)
	}
	err = ctx.DefaultDb.Write().Exec("INSERT INTO `user_manager` (`id`, `user_id`, `role`, `status`, `operator_id`, `ctime`, `mtime`) VALUES (1, 781, 1, 1, 1674, 1704847684, 1704847684);").Error
	if err != nil {
		logx.Infof("init database error: %+v", err)
	}
	err = ctx.DefaultDb.Write().Exec("INSERT INTO `user_white` (`id`, `user_id`, `cluster_id`, `status`, `ctime`, `mtime`, `operator_id`) VALUES (1, 806, 2, 1, 1704853147, 1704853147, 1674);").Error
	if err != nil {
		logx.Infof("init database error: %+v", err)
	}
}
