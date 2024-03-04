# obstwiesen-server



# TODO
- Adding Cultivars (Sorten)
- Adding Tags to everything

# Config

## Database-Connection
### SQLITE
- PROVIDER = sqlite
- CON_STR = path to file (data.db)

### MySQL
- PROVIDER = mysql
- CON_STR = root:my-secret-pw@tcp(127.0.0.1:3306)/meadows


## File-Storage
- PUBLIC_URL = "http://localhost:8080"
- FILE_PROVIDER = "disk"
- FILE_CONNSTR = "files" 