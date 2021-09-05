ccache parser and exporter
==========================

A `ccache`_ parser and `Prometheus`_ exporter.

====== ======
Branch Status
====== ======
master .. image:: https://travis-ci.com/virtualtam/ccache_exporter.svg?branch=master
          :target: https://travis-ci.com/virtualtam/ccache_exporter
          :alt: Travis build status
====== ======


Metrics exposed
---------------

Counters (internal):

- ``ccache_collector_parsing_errors_total``

Counters (ccache):

- ``ccache_call_total``
- ``ccache_call_hit_total``
- ``ccache_called_for_link_total``
- ``ccache_called_for_preprocessing_total``
- ``ccache_unsupported_code_directive_total``
- ``ccache_no_input_file_total``
- ``ccache_cleanups_performed_total``


Gauges (ccache):

- ``ccache_cache_hit_ratio``
- ``ccache_cache_size_bytes``
- ``ccache_cache_size_max_bytes``
- ``ccache_cached_files``

Building
--------

::

    # get the sources
    $ git clone https://github.com/virtualtam/ccache_exporter.git

    # build the parser and exporter
    $ make build

Parser usage
------------

::

   $ ccache -s | ccacheparser | jq

   {
     "cache_directory": "/home/virtualtam/.ccache",
     "primary_config": "/home/virtualtam/.ccache/ccache.conf",
     "secondary_config_readonly": "/etc/ccache.conf",
     "stats_time": "2018-09-24T21:19:07.997866938+02:00",
     "stats_zero_time": "2018-09-23T01:18:52+02:00",
     "cache_hit_direct": 124,
     "cache_hit_preprocessed": 8,
     "cache_miss": 297,
     "cache_hit_rate": 30.77,
     "called_for_link": 39,
     "called_for_preprocessing": 263,
     "unsupported_code_directive": 5,
     "no_input_file": 83,
     "cleanups_performed": 0,
     "files_in_cache": 926,
     "cache_size": "17.5 MB",
     "cache_size_bytes": 17500000,
     "max_cache_size": "15.0 GB",
     "max_cache_size_bytes": 15000000000
   }

Running with Docker Compose
---------------------------

The provided ``docker-compose.yml`` script defines the following monitoring
stack:

- ``ccache-exporter`` service bind-mounted on the user's (hint: that's you \\o/)
  ``$HOME`` directory;
- ``node-exporter`` service to gather system metrics;
- ``prometheus`` database, preconfigured to scrap exported ``ccache`` metrics;
- ``grafana`` dashboard, preconfigured to use ``prometheus`` as a data source
  and display ``ccache`` metrics in the corresponding dashboard.


To pull Docker images and start the services:

::

    $ docker-compose pull
    $ docker-compose up -d


Once the stack is up, the following services will be exposed:

- http://localhost:19508: ``ccache-exporter``
- http://localhost:19100: ``node-exporter``
- http://localhost:19090: ``prometheus``
- http://localhost:13000: ``grafana``


Then, login to Grafana with the default credentials (``admin/admin``) and load
the ``ccache`` dashboard:


.. image:: dashboard.jpg


.. _ccache: https://ccache.samba.org/
.. _Prometheus: https://prometheus.io/
