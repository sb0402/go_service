require 'liquid'
require 'json'
require 'benchmark'
module LiquidParsing 
  def self.main
    html = File.read("/Users/0352mpind/Documents/projects/go_service/templates/dynamic_tags")
    data = File.read("/Users/0352mpind/Documents/projects/go_service/data.json")
    tags = JSON.parse(data)
    variables = {}
    variables["campaign"] = tags["data"]["campaign"]
    variables["contact"] = tags["data"]["contact"]
    variables["other"] = tags["data"]["other"]
    variables["dynamic_content"] = tags["data"]["dynamic_content"]
    variables["table"] = tags["data"]["table"]
    variables["email"] = tags["data"]["email"]
    variables["template"] = tags["data"]["template"]
    variables["coupon"] = tags["data"]["coupon"]
    puts Benchmark.measure{response = Liquid::Template.parse(html).render(variables)}
  end
end

LiquidParsing.main
