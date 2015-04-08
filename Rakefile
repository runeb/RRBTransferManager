desc 'test'
task :test do
  sh "xcodebuild -project RRBTransferManager.xcodeproj -scheme \"RRBTransferManager\" CONFIGURATION_BUILD_DIR='build' clean test"
end
task :default => :test
