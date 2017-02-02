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
  headers 'Access-Control-Allow-Origin' => 'http://localhost:3000'
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
  stddev_in = request.body.read
  stddev = JSON.parse(stddev_in)

  stddevnew = StandardDeviationTbl.new
  stddevnew.answer = stddev['points'].standard_deviation
  stddevnew.input_data = stddev_in
  stddevnew.save

  stddev['answer'] = stddevnew.answer
  stddev['id'] = stddevnew.id
  stddev.to_json
end

def db_decode_stddev(stddev)
  {:id => stddev.id, :answer => stddev.answer, :points => JSON.parse(stddev.input_data)}
end
