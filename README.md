导表：
docker exec -i mysql mysql -h 127.0.0.1 -P 3306 -u root -p'rootroot' < /home/github/ranger/dao/gozero_sys_dept.sql

清库：
docker exec -i mysql mysql -h 127.0.0.1 -P 3306 -u root -p'rootroot' -e "drop database gozero"

show variables like 'character%';