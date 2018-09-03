export my_hostname=$(cat etcd-install.json | jq -r '.hostname')
export my_ip=$(cat etcd-install.json | jq -r '.nodes[] | select(.hostname=="'$my_hostname'") | .ip')
export cluster_list=$(cat etcd-install.json | jq -r '.nodes[] | "\(.hostname)=https://\(.ip):2380"' | paste -s -d ',')

echo "Installing ETCD on  : $my_hostname"
echo "Setting ip to       : $my_ip"
echo "etcd cluster string : $cluster_list"
