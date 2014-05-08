Frontend for g0
==

g0 frontend is based on dartlang and scss. 

##Setup

We use grunt to compile dart to js and SCSS to CSS. Therfore the dart SDK and 
node.js are required. 

* Run **pub get** to install dart dependencies
* Run **npm install** to install grunt dev dependencies

##Dev
* Rename **config-sample.json** to **config.json** in *web/*
* Run **grunt watch** to build css when a scss file gets changed 

##Build

* Run **grunt build-test** to build test version
* Run **grunt build-live** to build compressed live version 

##Deploy

Rename **.deploy-sample.json** to **.deploy.json** 
and **.ftppass-sample** to **.ftppass** and edit your server settings

* Run **grunt deploy-test** to deploy to live server (uses build-test)
* Run **grunt deploy-live** to deploy to test server (uses build-live)