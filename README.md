shit
-----

S.H.I.T stands for Security Heuristic Information Tooling. SHIT is a static analysis tool that scans codebases for known security vulnerabilities. This is one of the most advance pieces of tooling available, revealing signifcantly more vulnerabilities than competing tools. The result of these scans is output in the Code Representing A Problem or CRAP reports. 


## Features

* `shit` can output its CRAP reports in either XML or CSV format, which can be specified by the `CRAP_OUTPUT_REPORT_FORMAT` environment variable
* `shit` can output its CRAP reports to a file specified by `CRAP_OUTPUT_REPORT_FILE` or to standard output (but in standard output mode it will be mixed with debug information).
* Does not return a nonzero exit code if violations are found, so as not to disrupt CI/CD. To use as a check, simply inspect the contents of the CRAP reports.
* Unlike popular open source tools, ignoring "false positives" is not possible. 
* Support for the top ten most popular languages on GitHub right out of the box!

## Example

```
$ shit
 --- Security Hueristic Information Tooling ---
 --- Checking Files ---
 checking:  .gitignore
 checking:  .rspec
 checking:  .ruby-version
 checking:  Gemfile
 checking:  Gemfile.lock
 checking:  README.md
 checking:  Rakefile
 checking:  app/assets/config/manifest.js
 checking:  app/assets/images/.keep
 checking:  app/assets/stylesheets/application.css
 checking:  app/channels/application_cable/channel.rb
 checking:  app/channels/application_cable/connection.rb
 checking:  app/controllers/application_controller.rb
 checking:  app/controllers/concerns/.keep
 checking:  app/helpers/application_helper.rb
 checking:  app/javascript/channels/consumer.js
 checking:  app/javascript/channels/index.js
 checking:  app/javascript/packs/application.js
 checking:  app/jobs/application_job.rb
 checking:  app/mailers/application_mailer.rb
 checking:  app/models/application_record.rb
 checking:  app/models/concerns/.keep
 checking:  app/views/layouts/application.html.erb
 checking:  app/views/layouts/mailer.html.erb
 checking:  app/views/layouts/mailer.text.erb
 checking:  bin/bundle
 checking:  bin/env_setup
 checking:  bin/rails
 checking:  bin/rake
 checking:  bin/setup
 checking:  bin/spring
 checking:  config.ru
 checking:  config/application.rb
 checking:  config/boot.rb
 checking:  config/cable.yml
 checking:  config/credentials.yml.enc
 checking:  config/database.yml
 checking:  config/environment.rb
 checking:  config/environments/development.rb
 checking:  config/environments/production.rb
 checking:  config/environments/test.rb
 checking:  config/initializers/application_controller_renderer.rb
 checking:  config/initializers/assets.rb
 checking:  config/initializers/backtrace_silencers.rb
 checking:  config/initializers/content_security_policy.rb
 checking:  config/initializers/cookies_serializer.rb
 checking:  config/initializers/cors.rb
 checking:  config/initializers/filter_parameter_logging.rb
 checking:  config/initializers/inflections.rb
 checking:  config/initializers/mime_types.rb
 checking:  config/initializers/wrap_parameters.rb
 checking:  config/locales/en.yml
 checking:  config/locales/es.yml
 checking:  config/master.key
 checking:  config/puma.rb
 checking:  config/routes.rb
 checking:  config/spring.rb
 checking:  config/storage.yml
 checking:  config/webpack/development.js
 checking:  config/webpack/environment.js
 checking:  config/webpack/production.js
 checking:  config/webpack/test.js
 checking:  config/webpacker.yml
 checking:  lib/tasks/.keep
 checking:  log/.keep
 checking:  public/robots.txt
 checking:  reports/.keep
 checking:  spec/rails_helper.rb
 checking:  spec/sample_spec.rb
 checking:  spec/spec_helper.rb
 checking:  storage/.keep
 checking:  tmp/.keep
 checking:  vendor/.keep
 --- Generating CRAP Report ---
 <CrapReport>
 <Violation Filename=".gitignore" Level="critical">SHIT doesn&#39;t recognize this file, which is suspicious!!!</Violation>
 <Violation Filename=".rspec" Level="critical">SHIT doesn&#39;t recognize this file, which is suspicious!!!</Violation>
 <Violation Filename=".ruby-version" Level="critical">SHIT doesn&#39;t recognize this file, which is suspicious!!!</Violation>
 <Violation Filename="Gemfile" Level="critical">SHIT doesn&#39;t recognize this file, which is suspicious!!!</Violation>
 <Violation Filename="Gemfile.lock" Level="critical">SHIT doesn&#39;t recognize this file, which is suspicious!!!</Violation>
 <Violation Filename="README.md" Level="critical">SHIT doesn&#39;t recognize this file, which is suspicious!!!</Violation>
 <Violation Filename="Rakefile" Level="critical">SHIT doesn&#39;t recognize this file, which is suspicious!!!</Violation>
 <Violation Filename="app/assets/config/manifest.js" Level="criticaler">JavaScript is a programming language that can be used to execute malicious code!!!</Violation>
 <Violation Filename="app/assets/images/.keep" Level="critical">SHIT doesn&#39;t recognize this file, which is suspicious!!!</Violation>
 <Violation Filename="app/assets/stylesheets/application.css" Level="critical">SHIT doesn&#39;t recognize this file, which is suspicious!!!</Violation>
 <Violation Filename="app/channels/application_cable/channel.rb" Level="criticaler">Ruby is a programming language that can be used to execute malicious code!!!</Violation>
 <Violation Filename="app/channels/application_cable/connection.rb" Level="criticaler">Ruby is a programming language that can be used to execute malicious code!!!</Violation>
 <Violation Filename="app/controllers/application_controller.rb" Level="criticaler">Ruby is a programming language that can be used to execute malicious code!!!</Violation>
 <Violation Filename="app/controllers/concerns/.keep" Level="critical">SHIT doesn&#39;t recognize this file, which is suspicious!!!</Violation>
 <Violation Filename="app/helpers/application_helper.rb" Level="criticaler">Ruby is a programming language that can be used to execute malicious code!!!</Violation>
 <Violation Filename="app/javascript/channels/consumer.js" Level="criticaler">JavaScript is a programming language that can be used to execute malicious code!!!</Violation>
 <Violation Filename="app/javascript/channels/index.js" Level="criticaler">JavaScript is a programming language that can be used to execute malicious code!!!</Violation>
 <Violation Filename="app/javascript/packs/application.js" Level="criticaler">JavaScript is a programming language that can be used to execute malicious code!!!</Violation>
 <Violation Filename="app/jobs/application_job.rb" Level="criticaler">Ruby is a programming language that can be used to execute malicious code!!!</Violation>
 <Violation Filename="app/mailers/application_mailer.rb" Level="criticaler">Ruby is a programming language that can be used to execute malicious code!!!</Violation>
 <Violation Filename="app/models/application_record.rb" Level="criticaler">Ruby is a programming language that can be used to execute malicious code!!!</Violation>
 <Violation Filename="app/models/concerns/.keep" Level="critical">SHIT doesn&#39;t recognize this file, which is suspicious!!!</Violation>
 <Violation Filename="app/views/layouts/application.html.erb" Level="critical">SHIT doesn&#39;t recognize this file, which is suspicious!!!</Violation>
 <Violation Filename="app/views/layouts/mailer.html.erb" Level="critical">SHIT doesn&#39;t recognize this file, which is suspicious!!!</Violation>
 <Violation Filename="app/views/layouts/mailer.text.erb" Level="critical">SHIT doesn&#39;t recognize this file, which is suspicious!!!</Violation>
 <Violation Filename="bin/bundle" Level="criticalest">executable files are inherently dangerous!!!</Violation>
 <Violation Filename="bin/env_setup" Level="criticalest">executable files are inherently dangerous!!!</Violation>
 <Violation Filename="bin/rails" Level="criticalest">executable files are inherently dangerous!!!</Violation>
 <Violation Filename="bin/rake" Level="criticalest">executable files are inherently dangerous!!!</Violation>
 <Violation Filename="bin/setup" Level="criticalest">executable files are inherently dangerous!!!</Violation>
 <Violation Filename="bin/spring" Level="criticalest">executable files are inherently dangerous!!!</Violation>
 <Violation Filename="config.ru" Level="criticaler">Ruby is a programming language that can be used to execute malicious code!!!</Violation>
 <Violation Filename="config/application.rb" Level="criticaler">Ruby is a programming language that can be used to execute malicious code!!!</Violation>
 <Violation Filename="config/boot.rb" Level="criticaler">Ruby is a programming language that can be used to execute malicious code!!!</Violation>
 <Violation Filename="config/cable.yml" Level="critical">SHIT doesn&#39;t recognize this file, which is suspicious!!!</Violation>
 <Violation Filename="config/credentials.yml.enc" Level="critical">SHIT doesn&#39;t recognize this file, which is suspicious!!!</Violation>
 <Violation Filename="config/database.yml" Level="critical">SHIT doesn&#39;t recognize this file, which is suspicious!!!</Violation>
 <Violation Filename="config/environment.rb" Level="criticaler">Ruby is a programming language that can be used to execute malicious code!!!</Violation>
 <Violation Filename="config/environments/development.rb" Level="criticaler">Ruby is a programming language that can be used to execute malicious code!!!</Violation>
 <Violation Filename="config/environments/production.rb" Level="criticaler">Ruby is a programming language that can be used to execute malicious code!!!</Violation>
 <Violation Filename="config/environments/test.rb" Level="criticaler">Ruby is a programming language that can be used to execute malicious code!!!</Violation>
 <Violation Filename="config/initializers/application_controller_renderer.rb" Level="criticaler">Ruby is a programming language that can be used to execute malicious code!!!</Violation>
 <Violation Filename="config/initializers/assets.rb" Level="criticaler">Ruby is a programming language that can be used to execute malicious code!!!</Violation>
 <Violation Filename="config/initializers/backtrace_silencers.rb" Level="criticaler">Ruby is a programming language that can be used to execute malicious code!!!</Violation>
 <Violation Filename="config/initializers/content_security_policy.rb" Level="criticaler">Ruby is a programming language that can be used to execute malicious code!!!</Violation>
 <Violation Filename="config/initializers/cookies_serializer.rb" Level="criticaler">Ruby is a programming language that can be used to execute malicious code!!!</Violation>
 <Violation Filename="config/initializers/cors.rb" Level="criticaler">Ruby is a programming language that can be used to execute malicious code!!!</Violation>
 <Violation Filename="config/initializers/filter_parameter_logging.rb" Level="criticaler">Ruby is a programming language that can be used to execute malicious code!!!</Violation>
 <Violation Filename="config/initializers/inflections.rb" Level="criticaler">Ruby is a programming language that can be used to execute malicious code!!!</Violation>
 <Violation Filename="config/initializers/mime_types.rb" Level="criticaler">Ruby is a programming language that can be used to execute malicious code!!!</Violation>
 <Violation Filename="config/initializers/wrap_parameters.rb" Level="criticaler">Ruby is a programming language that can be used to execute malicious code!!!</Violation>
 <Violation Filename="config/locales/en.yml" Level="critical">SHIT doesn&#39;t recognize this file, which is suspicious!!!</Violation>
 <Violation Filename="config/locales/es.yml" Level="critical">SHIT doesn&#39;t recognize this file, which is suspicious!!!</Violation>
 <Violation Filename="config/master.key" Level="critical">SHIT doesn&#39;t recognize this file, which is suspicious!!!</Violation>
 <Violation Filename="config/puma.rb" Level="criticaler">Ruby is a programming language that can be used to execute malicious code!!!</Violation>
 <Violation Filename="config/routes.rb" Level="criticaler">Ruby is a programming language that can be used to execute malicious code!!!</Violation>
 <Violation Filename="config/spring.rb" Level="criticaler">Ruby is a programming language that can be used to execute malicious code!!!</Violation>
 <Violation Filename="config/storage.yml" Level="critical">SHIT doesn&#39;t recognize this file, which is suspicious!!!</Violation>
 <Violation Filename="config/webpack/development.js" Level="criticaler">JavaScript is a programming language that can be used to execute malicious code!!!</Violation>
 <Violation Filename="config/webpack/environment.js" Level="criticaler">JavaScript is a programming language that can be used to execute malicious code!!!</Violation>
 <Violation Filename="config/webpack/production.js" Level="criticaler">JavaScript is a programming language that can be used to execute malicious code!!!</Violation>
 <Violation Filename="config/webpack/test.js" Level="criticaler">JavaScript is a programming language that can be used to execute malicious code!!!</Violation>
 <Violation Filename="config/webpacker.yml" Level="critical">SHIT doesn&#39;t recognize this file, which is suspicious!!!</Violation>
 <Violation Filename="lib/tasks/.keep" Level="critical">SHIT doesn&#39;t recognize this file, which is suspicious!!!</Violation>
 <Violation Filename="log/.keep" Level="critical">SHIT doesn&#39;t recognize this file, which is suspicious!!!</Violation>
 <Violation Filename="public/robots.txt" Level="critical">SHIT doesn&#39;t recognize this file, which is suspicious!!!</Violation>
 <Violation Filename="reports/.keep" Level="critical">SHIT doesn&#39;t recognize this file, which is suspicious!!!</Violation>
 <Violation Filename="spec/rails_helper.rb" Level="criticaler">Ruby is a programming language that can be used to execute malicious code!!!</Violation>
 <Violation Filename="spec/sample_spec.rb" Level="criticaler">Ruby is a programming language that can be used to execute malicious code!!!</Violation>
 <Violation Filename="spec/spec_helper.rb" Level="criticaler">Ruby is a programming language that can be used to execute malicious code!!!</Violation>
 <Violation Filename="storage/.keep" Level="critical">SHIT doesn&#39;t recognize this file, which is suspicious!!!</Violation>
 <Violation Filename="tmp/.keep" Level="critical">SHIT doesn&#39;t recognize this file, which is suspicious!!!</Violation>
 <Violation Filename="vendor/.keep" Level="critical">SHIT doesn&#39;t recognize this file, which is suspicious!!!</Violation>
```

## Future Development

* More language support
* Documentation on how to actually run this (low priority)
* XML could be made more arcane and unreadable by humans
* Docker image

## Installation

Please don't
