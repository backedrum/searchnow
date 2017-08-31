## searchnow command line search tool
*Searchnow* allows you to perform a quick web search or get instant info from various APIs directly from the command line. 

### Make
> make

### Use searchnow:
> searchnow <search_term> <search_engine> <max_results>

### config.json
*config.json* should be placed into the app directory. Use it to include/exclude particular fields from result.

### Supported engines:
- Google. Make sure you comply with https://console.developers.google.com/tos?id=universal
  
  This project utilizes [Google Custom Search JSON/Atom API](https://developers.google.com/custom-search/json-api/v1/overview)
  
  Place your tokens into the .tokens file in the app directory. You do need to specify google.apikey - it is your personal app key to use Custom Search API.
  
  Also you do need to create a specific engine inside of the Google Control Panel and to provide engine id in google.engine token.
  So your .tokens file will be look like this:
  ````
  google.apikey=<your_api_key_associated_with_custom_search_api>
  google.engine=<search_engine_id>
  ````
  Command line value: google.
- StackOverflow. Performs search for the questions from stackoverflow.com via [Stack-on-Go](https://github.com/laktek/Stack-on-Go).
  Command line value: so.
- Hacker News. Allows fetching of the stories from [Hacker News](news.ycombinator.com) via [Hacker News Api](https://github.com/HackerNews/API).
  Supported terms: *newstories, topstories, beststories, askstories, jobstories*.
  Command line value: hn.
- ipvigilante.com. Retrieves location by IP via [ipvigilante.com Api](www.ipvigilante.com).
  Command line value: ip_loc.     
- Other engines (might be added soon).

### Command line arguments:
- *search_term* (required)     Term to search for search engines or action/expression for other engines.
- *search_engine* (optional)   Search engine, "google" is default one.
- *max_resuts* (optional)      Limit a number of results. Default value is 5.  

### Samples:
> searchnow "go: unknown subcommand" google 10  
> searchnow "factorial" so 9  
> searchnow "newstories" hn  
> searchnow "8.8.8.8" ip_loc  

### License:
> Apache 2.0
                                        
