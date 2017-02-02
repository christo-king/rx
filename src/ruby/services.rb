require 'sinatra'
require 'json'
require 'active_record'
load 'stddev.rb'

ActiveRecord::Base.establish_connection ({
    :adapter => "mysql2",
    :host => "rxdb",
    :database => "testing",
    :username => "test_user",
    :password => "test1"
})

class StandardDeviationTbl < ActiveRecord::Base
  self.table_name = "standard_deviation_tbl"
end

before do
  request.body.rewind
  @request_payload = JSON.parse request.body.read
end

get '/standardDeviation' do
  content_type :json
  std_devs = []
  StandardDeviationTbl.find_each do |stddev|
    std_devs << db_decode_stddev(stddev)
  end
  std_devs.to_json
end

get '/standardDeviation/:id' do |id|
  content_type :json
  db_decode_stddev(StandardDeviationTbl.find(id)).to_json
end

post '/standardDeviation' do
  stddev = JSON.parse(request.body.read)
  stddevnew = StandardDeviationTbl.new do |sd|
    sd.answer = stddev.points.standard_deviation
  end
end

def db_decode_stddev(stddev)
  {:id => stddev.id, :answer => stddev.answer, :points => JSON.parse(stddev.input_data)}
end