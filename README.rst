ccache exporter
===============

A `Prometheus`_ exporter that exposes `ccache`_ metrics.

Metrics exposed
---------------

Counters:

- cache_hit_total
- cache_miss_total
- called_for_link_total
- called_for_preprocessing_total
- unsupported_code_directive_total
- no_input_file_total
- cleanups_performed_total


Gauges:

- cache_hit_ratio
- cache_size_bytes
- max_cache_size_bytes
- files_in_cache


Running with Docker Compose
---------------------------

The provided ``docker-compose.yml`` script defines the following monitoring
stack:

- ``ccache-exporter`` service bind-mounted on the user's (hint: that's you \o/)
  ``$HOME`` directory;
- ``prometheus`` database, preconfigured to scrap exported ``ccache`` metrics;
- ``grafana`` dashboard, preconfigured to use ``prometheus`` as a data source
  and display ``ccache`` metrics in the corresponding dashboard.


To pull Docker images and start the services:

::

    $ docker-compose pull
    $ docker-compose up -d


Once the stack is up, the following ports will be exposed:

- ``19508``: ``ccache-exporter``
- ``19090``: ``prometheus``
- ``13000``: ``grafana``


Then, login to Grafana with the default credentials (``admin/admin``) and load
the ``ccache`` dashboard:


.. image:: dashboard.jpg


.. _ccache: https://ccache.samba.org/
.. _Prometheus: https://prometheus.io/
