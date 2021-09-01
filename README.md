# Open Stocks Exporter

Prometheus exporter for stocks. Open-Stocks-Exporter uses Yahoo finance API
under the hood, to fetch stock entities in real-time and expose them as Prometheus metrics.

Metrics are designed with analytics in mind. This helps answer any questions, take any actions or decisions related to stock market. You can write advanced analytical
queries in SQL (Postgres) if you use [Promscale](https://github.com/timescale/promscale) as remote-storage for your Prometheus.

Based on some quick tests, the values are updated with no delay.

Some examples of queries will be added in the later part of development. Stay tuned!


## Contributions

Contributions are welcome. Please open issues for bug reports or feature requests.
