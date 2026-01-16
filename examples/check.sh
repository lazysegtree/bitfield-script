while true; do
    for host in $(cat servers.txt); do
      echo "[$(date)] $host -> $(dig +short $host | head -n 1)"
    done
  sleep 1
done