total_requests = 200
concurrency = 10
url = "https://httpbin.org/anything"
method = :post
payload = '{"some": "data"}'

:inets.start()
:ssl.start()

average_time =
  1..total_requests
  |> Task.async_stream(
    fn _ ->
      start_time = System.monotonic_time()
      :httpc.request(method, {url, [], 'application/json', payload}, [], [])
      System.monotonic_time() - start_time
    end,
    max_concurrency: concurrency
  )
  |> Enum.reduce(0, fn {:ok, req_time}, acc ->
    acc + req_time
  end)
  |> System.convert_time_unit(:native, :millisecond)
  |> Kernel./(total_requests)

IO.puts("Average response time from #{url} is #{average_time}ms")

# $ time elixir elixir_stress.exs
# Average response time from https://httpbin.org/anything is 0.195ms
# elixir elixir_stress.exs  0.49s user 0.27s system 140% cpu 0.538 total
