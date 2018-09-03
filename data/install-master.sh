

export my_hostname=$(cat master-install.json | jq -r '.hostname')
export my_ip=$(cat master-install.json | jq -r '.nodes[] | select(.hostname=="'$my_hostname'") | .ip')
export master_count=$(cat master-install.json | jq -r '.nodes | length')
export vip=$(cat master-install.json | jq -r '.vip')
export etcd_cluster=$(cat master-install.json | jq -r '.etcdnodes[] | "\(.hostname)=https://\(.ip):2380"' | paste -s -d ',')

echo "Installing Master on : $my_hostname"
echo "Setting ip to        : $my_ip"
echo "Number of masters    : $master_count"
echo "Vitual IP            : $vip"
echo "Etcd cluster string  : $etcd_cluster"
