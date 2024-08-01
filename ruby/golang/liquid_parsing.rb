require 'grpc'
require_relative 'liquid_pb'
require_relative 'liquid_services_pb'
require 'json'
require 'benchmark'

module LiquidParsing 
  def self.main
    stub = Liquid::LiquidParsing::Stub.new('127.0.0.1:50051', :this_channel_is_insecure)
    template = Liquid::Template.new
    html = File.read("/Users/0352mpind/Documents/projects/go_service/templates/dynamic_tags")
    data = File.read("/Users/0352mpind/Documents/projects/go_service/data.json")
    tags = JSON.parse(data)
    template.variables["html_part"] = html
    template.variables["campaign"] = JSON.generate(tags["data"]["campaign"])
    template.variables["contact"] = JSON.generate(tags["data"]["contact"])
    template.variables["other"] = JSON.generate(tags["data"]["other"])
    template.variables["dynamic_content"] = JSON.generate(tags["data"]["dynamic_content"])
    template.variables["table"] = JSON.generate(tags["data"]["table"])
    template.variables["email"] = JSON.generate(tags["data"]["email"])
    template.variables["template"] = JSON.generate(tags["data"]["template"])
    template.variables["coupon"] = JSON.generate(tags["data"]["coupon"])
    puts Benchmark.measure{response = stub.parse_and_render_string(template)}
    # begin
    #   response = stub.parse_and_render_string(template)
    #     puts "Response: #{response.result}"
    # rescue GRPC::BadStatus => e
    #   puts "Error: #{e.to_s}"
    # end
  end
end

LiquidParsing.main
