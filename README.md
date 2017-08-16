## searchnow command line search tool
*Searchnow* allows you to perform a quick web search or get instant info directly from the command line. 

### Make
> make

### Use searchnow:
> searchnow <search_term> <search_engine> <max_results>

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
- Other engines to be added soon.

### Command line arguments:
- *search_engine* (optional)   Search engine, "google" is default one.
- *max_resuts* (optional)      Limit a number of results. Default value is 5.  

### Sample:
> searchnow "go: unknown subcommand" google 10
> searchnow "factorial" so 9

### License:
> Apache 2.0
                                        
