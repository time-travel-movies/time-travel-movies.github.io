package main

import (
	"bufio"
	"errors"
	"html/template"
	"io"
	"log"
	"os"
	"strings"
)

var (
	errTSV = errors.New("need 5 values on a line, found less")
)

const indexTemplate = `<!DOCTYPE html>
<html lang="en">
	<head>
	    <meta charset="utf-8">
	    <meta http-equiv="X-UA-Compatible" content="IE=edge">
	    <meta name="viewport" content="width=device-width, initial-scale=1">

	    <meta name="description" content="A comprehensive list of movies that include time travel">
	    <meta name="author" content="The Time Travel Movie Consortium">

		<link rel="apple-touch-icon" sizes="180x180" href="/favicons/apple-touch-icon.png">
		<link rel="icon" type="image/png" sizes="32x32" href="/favicons/favicon-32x32.png">
		<link rel="icon" type="image/png" sizes="16x16" href="/favicons/favicon-16x16.png">
		<link rel="manifest" href="/favicons/manifest.json">
		<link rel="mask-icon" href="/favicons/safari-pinned-tab.svg" color="#5bbad5">
		<meta name="theme-color" content="#ffffff">
		 
		<meta name="application-name" content="Time Travel Movies"/>

		<meta property="og:title" content="Time Travel Movies" />
		<meta property="og:description" content="A comprehensive list of movies that include time travel" />	

		<title>Time Travel Movies</title>

		<link rel="stylesheet" type="text/css" href="/css/timetravelmovies.css">
		<link rel="stylesheet" type="text/css" href="/css/bootstrap.min.css">
	</head>

	<body>
	    
		<div class="container-fluid" style="background-color: #25062E;">
	      	<div class="row">
	      		<div class="col-md-12" style="display: flex; flex-flow: row; justify-content: center;">
	      			<img src="/images/time-travel-movies-logo.jpg" style="max-width: 750px; width: 100%; height: 100%;">

	      		</div>
	      	</div>
	    </div>

		<div class="container">
		
		{{range .}}
	    	
			<div class="row time-row">
				<div class="col-md-12">
					<div class="card-horizontal">
  						<a href="{{.IMDBLink}}"><img class="card-img-left" src="{{.Image}}"></a>
  						<div class="card-body-right">
    						<div class="card-text">
    							<a href="{{.IMDBLink}}"><h3>{{.Title}}</h3></a>
    							<h4>{{.Year}}</h4>
    							<p>{{.Description}}</p>
    						</div>
  						</div>
  					</div>
				</div>
	    	</div>    

		{{end}}

	    </div>

    	<script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.3/jquery.min.js"></script>
    	<script src="/css/bootstrap.min.js"></script>
    	
    	<script>
		  (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
		  (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
			m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
  			})(window,document,'script','https://www.google-analytics.com/analytics.js','ga');

  			ga('create', 'UA-106121766-1', 'auto');
  			ga('send', 'pageview');

		</script>

	</body>
</html>`

func main() {
	ee, err := parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	t, err := template.New("index").Parse(indexTemplate)
	if err != nil {
		log.Fatal(err)
	}
	if err := t.Execute(os.Stdout, ee); err != nil {
		log.Fatal(err)
	}
}

type entry struct {
	Title       string
	Year        string
	Image       string
	IMDBLink    string
	Description string
}

func parse(r io.Reader) ([]*entry, error) {
	var ee []*entry
	s := bufio.NewScanner(r)
	for s.Scan() {
		e, err := parseEntry(s.Text())
		if err != nil {
			return nil, err
		}
		ee = append(ee, e)
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return ee, nil
}

func parseEntry(s string) (*entry, error) {
	ss := strings.Split(s, "\t")
	if len(ss) < 5 {
		return nil, errTSV
	}
	e := &entry{
		Title:       ss[0],
		Year:        ss[1],
		Image:       ss[2],
		IMDBLink:    ss[3],
		Description: ss[4],
	}
	return e, nil
}
