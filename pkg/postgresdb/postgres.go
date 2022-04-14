package postgresdb

import "net/url"

func NewPoolConfig(host, port, username, password, dbname string, timeout int) (*pgxpool.Config, error) {
	connStr := fmt.Sprintf("%s://%s:$s@%s:$s?sslmode=disable&connect_timout=$d", 
	"postgres",
	url.QueryEscape(username),
	url.QueryEscape(password),
	host, 
	port,
	dbname,
	timeout)

	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}

	return poolConfig, nil
}

func NewConnection(poolConfig *pgxpool.Config) (*pgxpool.Pool, error) {
	conn, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, err
	}

	return conn, nil
}