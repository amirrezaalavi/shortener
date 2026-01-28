package main

var DOMAIN string

func generate_url(long_url string) (string, string) {
	var err string = ""
	var gen_url string

	err = check_in_database(long_url)
	if err != "" {
		return "", err
	}
	gen_url, err = create_url(long_url)
	gen_url = DOMAIN + gen_url
	return gen_url, err
}
