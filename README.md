# ONLINE SHOP PROJECTS
1. go mod init <nama project>
2. Jalankan docker untuk PostgreSQL database
```
docker run --name postgresql -e POSTGRES_USER=mastama -e POSTGRES_PASSWORD=post456 -e POSTGRES_DB=postgres -d -p 5432:5432 postgres:16
```