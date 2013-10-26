namespace :dependencies do
  desc "go get all the libraries in 'dependencies.txt'"
  task :install do
    deps = File.read('./dependencies.txt')
    deps.each_line do |l|
      cmd = "go get #{l}"
      puts cmd
      system(cmd)
    end
  end
end
