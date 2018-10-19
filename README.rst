ccache exporter
===============

A `Prometheus`_ exporter that exposes `ccache`_ metrics.


Metrics exposed
---------------

Counters:

- ``ccache_call_total``
- ``ccache_call_hit_total``
- ``ccache_called_for_link_total``
- ``ccache_called_for_preprocessing_total``
- ``ccache_unsupported_code_directive_total``
- ``ccache_no_input_file_total``
- ``ccache_cleanups_performed_total``


Gauges:

- ``ccache_cache_hit_ratio``
- ``ccache_cache_size_bytes``
- ``ccache_cache_size_max_bytes``
- ``ccache_cached_files``

Building
--------

::

    # get the sources
    $ git clone https://github.com/virtualtam/ccache_exporter.git

    # go >= 1.11: enable module support
    $ export GO111MODULE=on

    # build ccache_exporter
    $ make build


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
