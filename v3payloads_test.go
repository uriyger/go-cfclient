package cfclient

const listTasksByAppPayloadPage1 string = `
{
  "pagination": {
    "total_results": 4,
    "total_pages": 2,
    "first": {
      "href": "https://api.run.example.com/v3/apps/ccc25a0f-c8f4-4b39-9f1b-de9f328d0ee5/tasks?page=1&per_page=2"
    },
    "last": {
      "href": "https://api.run.example.com/v3/apps/ccc25a0f-c8f4-4b39-9f1b-de9f328d0ee5/tasks?page=2&per_page=2"
    },
    "next": {
      "href": "https://api.run.example.com/v3/apps/ccc25a0f-c8f4-4b39-9f1b-de9f328d0ee5/tasks?page=2&per_page=2"
    },
    "previous": null
  },
  "resources": [
  {
    "guid": "d5cc22ec-99a3-4e6a-af91-a44b4ab7b6fa",
    "sequence_id": 1,
    "name": "hello",
    "state": "SUCCEEDED",
    "memory_in_mb": 512,
    "disk_in_mb": 1024,
    "result": {
      "failure_reason": null
    },
    "droplet_guid": "740ebd2b-162b-469a-bd72-3edb96fabd9a",
    "metadata": {
      "labels": { },
      "annotations": { }
    },
    "created_at": "2016-05-04T17:00:41Z",
    "updated_at": "2016-05-04T17:00:42Z",
    "relationships": {
      "app": {
        "data": {
          "guid": "ccc25a0f-c8f4-4b39-9f1b-de9f328d0ee5"
        }
      }
    },
    "links": {
      "self": {
        "href": "https://api.example.org/v3/tasks/d5cc22ec-99a3-4e6a-af91-a44b4ab7b6fa"
      },
      "app": {
        "href": "https://api.example.org/v3/apps/ccc25a0f-c8f4-4b39-9f1b-de9f328d0ee5"
      },
      "cancel": {
        "href": "https://api.example.org/v3/tasks/d5cc22ec-99a3-4e6a-af91-a44b4ab7b6fa/actions/cancel",
        "method": "POST"
      },
      "droplet": {
        "href": "https://api.example.org/v3/droplets/740ebd2b-162b-469a-bd72-3edb96fabd9a"
      }
    }
  },
  {
    "guid": "63b4cd89-fd8b-4bf1-a311-7174fcc907d6",
    "sequence_id": 2,
    "name": "migrate",
    "state": "FAILED",
    "memory_in_mb": 1024,
    "disk_in_mb": 1024,
    "result": {
      "failure_reason": "Exited with status 1"
    },
    "droplet_guid": "740ebd2b-162b-469a-bd72-3edb96fabd9a",
    "metadata": {
      "labels": { },
      "annotations": { }
    },
    "created_at": "2016-05-04T17:00:43Z",
    "updated_at": "2016-05-04T17:00:44Z",
    "relationships": {
      "app": {
        "data": {
          "guid": "ccc25a0f-c8f4-4b39-9f1b-de9f328d0ee5"
        }
      }
    },
    "links": {
      "self": {
        "href": "https://api.example.org/v3/tasks/63b4cd89-fd8b-4bf1-a311-7174fcc907d6"
      },
      "app": {
        "href": "https://api.example.org/v3/apps/ccc25a0f-c8f4-4b39-9f1b-de9f328d0ee5"
      },
      "cancel": {
        "href": "https://api.example.org/v3/tasks/63b4cd89-fd8b-4bf1-a311-7174fcc907d6/actions/cancel",
        "method": "POST"
      },
      "droplet": {
        "href": "https://api.example.org/v3/droplets/740ebd2b-162b-469a-bd72-3edb96fabd9a"
      }
    }
  }
  ]
}
`

const listTasksByAppPayloadPage2 string = `
{
  "pagination": {
    "total_results": 4,
    "total_pages": 2,
    "first": {
      "href": "https://api.run.example.com/v3/apps/ccc25a0f-c8f4-4b39-9f1b-de9f328d0ee5/tasks?page=1&per_page=2"
    },
    "last": {
      "href": "https://api.run.example.com/v3/apps/ccc25a0f-c8f4-4b39-9f1b-de9f328d0ee5/tasks?page=2&per_page=2"
    },
    "next": null,
    "previous":{
      "href": "https://api.run.example.com/v3/apps/ccc25a0f-c8f4-4b39-9f1b-de9f328d0ee5/tasks?page=1&per_page=2"
    }
  },
  "resources": [
  {
    "guid": "abcdefc-99a3-4e6a-af91-a44b4ab7b6fa",
    "sequence_id": 3,
    "name": "hi",
    "state": "SUCCEEDED",
    "memory_in_mb": 512,
    "disk_in_mb": 1024,
    "result": {
      "failure_reason": null
    },
    "droplet_guid": "740ebd2b-162b-469a-bd72-3edb96fabd9a",
    "metadata": {
      "labels": { },
      "annotations": { }
    },
    "created_at": "2016-05-04T17:00:44Z",
    "updated_at": "2016-05-04T17:00:45Z",
    "relationships": {
      "app": {
        "data": {
          "guid": "ccc25a0f-c8f4-4b39-9f1b-de9f328d0ee5"
        }
      }
    },
    "links": {
      "self": {
        "href": "https://api.example.org/v3/tasks/abcdefc-99a3-4e6a-af91-a44b4ab7b6fa"
      },
      "app": {
        "href": "https://api.example.org/v3/apps/ccc25a0f-c8f4-4b39-9f1b-de9f328d0ee5"
      },
      "cancel": {
        "href": "https://api.example.org/v3/tasks/abcdefc-99a3-4e6a-af91-a44b4ab7b6fa/actions/cancel",
        "method": "POST"
      },
      "droplet": {
        "href": "https://api.example.org/v3/droplets/740ebd2b-162b-469a-bd72-3edb96fabd9a"
      }
    }
  },
  {
    "guid": "hijklm9-fd8b-4bf1-a311-7174fcc907d6",
    "sequence_id": 4,
    "name": "hello2",
    "state": "SUCCEEDED",
    "memory_in_mb": 1024,
    "disk_in_mb": 1024,
    "result": {
      "failure_reason": "Exited with status 1"
    },
    "droplet_guid": "740ebd2b-162b-469a-bd72-3edb96fabd9a",
    "metadata": {
      "labels": { },
      "annotations": { }
    },
    "created_at": "2016-05-04T17:00:46Z",
    "updated_at": "2016-05-04T17:00:47Z",
    "relationships": {
      "app": {
        "data": {
          "guid": "ccc25a0f-c8f4-4b39-9f1b-de9f328d0ee5"
        }
      }
    },
    "links": {
      "self": {
        "href": "https://api.example.org/v3/tasks/hijklm9-fd8b-4bf1-a311-7174fcc907d6"
      },
      "app": {
        "href": "https://api.example.org/v3/apps/ccc25a0f-c8f4-4b39-9f1b-de9f328d0ee5"
      },
      "cancel": {
        "href": "https://api.example.org/v3/tasks/hijklm9-fd8b-4bf1-a311-7174fcc907d6/actions/cancel",
        "method": "POST"
      },
      "droplet": {
        "href": "https://api.example.org/v3/droplets/740ebd2b-162b-469a-bd72-3edb96fabd9a"
      }
    }
  }
  ]
}
`

const listTasksPayloadPage1 string = `
{
  "pagination": {
    "total_results": 4,
    "total_pages": 2,
    "first": {
      "href ": "https: //api.run.example.com/v3/tasks?page=1&per_page=2"
    },
    "last": {
      "href": "https://api.run.example.com/v3/tasks?page=2&per_page=2"
    },
    "next": {
      "href": "https://api.run.example.com/v3/tasks?page=2&per_page=2"
    },
    "previous": null
  },
  "resources": [{
      "guid": "d5cc22ec-99a3-4e6a-af91-a44b4ab7b6fa",
      "sequence_id": 1,
      "name": "hello",
      "state": "SUCCEEDED",
      "memory_in_mb": 512,
      "disk_in_mb": 1024,
      "result": {
        "failure_reason": null
      },
      "droplet_guid": "740ebd2b-162b-469a-bd72-3edb96fabd9a",
      "metadata": {
        "labels": {},
        "annotations": {}
      },
      "created_at": "2016-05-04T17:00:41Z",
      "updated_at": "2016-05-04T17:00:42Z",
      "links": {
        "self": {
          "href": "https://api.example.org/v3/tasks/d5cc22ec-99a3-4e6a-af91-a44b4ab7b6fa"
        },
        "app": {
          "href": "https://api.example.org/v3/apps/ccc25a0f-c8f4-4b39-9f1b-de9f328d0ee5"
        },
        "cancel": {
          "href": "https://api.example.org/v3/tasks/d5cc22ec-99a3-4e6a-af91-a44b4ab7b6fa/actions/cancel",
          "method": "POST"
        },
        "droplet": {
          "href": "https://api.example.org/v3/droplets/740ebd2b-162b-469a-bd72-3edb96fabd9a"
        }
      }
    },
    {
      "guid": "63b4cd89-fd8b-4bf1-a311-7174fcc907d6",
      "sequence_id": 2,
      "name": "migrate",
      "state": "FAILED",
      "memory_in_mb": 1024,
      "disk_in_mb": 1024,
      "result": {
        "failure_reason": "Exited with status 1"
      },
      "droplet_guid": "740ebd2b-162b-469a-bd72-3edb96fabd9a",
      "metadata": {
        "labels": {},
        "annotations": {}
      },
      "created_at": "2016-05-04T17:00:43Z",
      "updated_at": "2016-05-04T17:00:44Z",
      "links": {
        "self": {
          "href": "https://api.example.org/v3/tasks/63b4cd89-fd8b-4bf1-a311-7174fcc907d6"
        },
        "app": {
          "href": "https://api.example.org/v3/apps/ccc25a0f-c8f4-4b39-9f1b-de9f328d0ee5"
        },
        "cancel": {
          "href": "https://api.example.org/v3/tasks/63b4cd89-fd8b-4bf1-a311-7174fcc907d6/actions/cancel",
          "method": "POST"
        },
        "droplet": {
          "href": "https://api.example.org/v3/droplets/740ebd2b-162b-469a-bd72-3edb96fabd9a"
        }
      }
    }
  ]
}`

const listTasksPayloadPage2 string = `
{
  "pagination": {
    "total_results": 4,
    "total_pages": 2,
    "first": {
      "href": "https://api.run.example.com/v3/tasks?page=1&per_page=2"
    },
    "last": {
      "href": "https://api.run.example.com/v3/tasks?page=2&per_page=2"
    },
    "next": null,
    "previous": {
      "href": "https://api.run.example.com/v3/tasks?page=1&per_page=2"
    }
  },
  "resources": [
  {
    "guid": "abcdefc-99a3-4e6a-af91-a44b4ab7b6fa",
    "sequence_id": 3,
    "name": "hi",
    "state": "SUCCEEDED",
    "memory_in_mb": 512,
    "disk_in_mb": 1024,
    "result": {
      "failure_reason": null
    },
    "droplet_guid": "740ebd2b-162b-469a-bd72-3edb96fabd9a",
    "metadata": {
      "labels": { },
      "annotations": { }
    },
    "created_at": "2016-05-04T17:00:44Z",
    "updated_at": "2016-05-04T17:00:45Z",
    "links": {
      "self": {
        "href": "https://api.example.org/v3/tasks/abcdefc-99a3-4e6a-af91-a44b4ab7b6fa"
      },
      "app": {
        "href": "https://api.example.org/v3/apps/ccc25a0f-c8f4-4b39-9f1b-de9f328d0ee5"
      },
      "cancel": {
        "href": "https://api.example.org/v3/tasks/abcdefc-99a3-4e6a-af91-a44b4ab7b6fa/actions/cancel",
        "method": "POST"
      },
      "droplet": {
        "href": "https://api.example.org/v3/droplets/740ebd2b-162b-469a-bd72-3edb96fabd9a"
      }
    }
  },
  {
    "guid": "hijklm9-fd8b-4bf1-a311-7174fcc907d6",
    "sequence_id": 4,
    "name": "hello2",
    "state": "SUCCEEDED",
    "memory_in_mb": 1024,
    "disk_in_mb": 1024,
    "result": {
      "failure_reason": "Exited with status 1"
    },
    "droplet_guid": "740ebd2b-162b-469a-bd72-3edb96fabd9a",
    "metadata": {
      "labels": { },
      "annotations": { }
    },
    "created_at": "2016-05-04T17:00:46Z",
    "updated_at": "2016-05-04T17:00:47Z",
    "links": {
      "self": {
        "href": "https://api.example.org/v3/tasks/hijklm9-fd8b-4bf1-a311-7174fcc907d6"
      },
      "app": {
        "href": "https://api.example.org/v3/apps/ccc25a0f-c8f4-4b39-9f1b-de9f328d0ee5"
      },
      "cancel": {
        "href": "https://api.example.org/v3/tasks/hijklm9-fd8b-4bf1-a311-7174fcc907d6/actions/cancel",
        "method": "POST"
      },
      "droplet": {
        "href": "https://api.example.org/v3/droplets/740ebd2b-162b-469a-bd72-3edb96fabd9a"
      }
    }
  }
  ]
}
`

const createTaskPayload = `
{
  "guid": "d5cc22ec-99a3-4e6a-af91-a44b4ab7b6fa",
  "sequence_id": 1,
  "name": "migrate",
  "command": "rake db:migrate",
  "state": "RUNNING",
  "memory_in_mb": 512,
  "disk_in_mb": 1024,
  "result": {
    "failure_reason": null
  },
  "droplet_guid": "740ebd2b-162b-469a-bd72-3edb96fabd9a",
  "created_at": "2016-05-04T17:00:41Z",
  "updated_at": "2016-05-04T17:00:42Z",
  "links": {
    "self": {
      "href": "https://api.example.org/v3/tasks/d5cc22ec-99a3-4e6a-af91-a44b4ab7b6fa"
    },
    "app": {
      "href": "https://api.example.org/v3/apps/ccc25a0f-c8f4-4b39-9f1b-de9f328d0ee5"
    },
    "droplet": {
      "href": "https://api.example.org/v3/droplets/740ebd2b-162b-469a-bd72-3edb96fabd9a"
    }
  }
}
`

const listProcessesPayload1 = `{
  "pagination": {
    "total_results": 26,
    "total_pages": 2,
    "first": {
      "href": "https://api.run.example.com/v3/processes?page=1&per_page=20"
    },
    "last": {
      "href": "https://api.run.example.com/v3/processes?page=2&per_page=20"
    },
    "next": {
      "href": "https://api.run.example.com/v3/processes?page=2&per_page=20"
    },
    "previous": null
  },
  "resources": [
    {
      "guid": "30200dd1-64c0-451a-8c37-3c91e05c6741",
      "type": "web",
      "command": "[PRIVATE DATA HIDDEN IN LISTS]",
      "instances": 2,
      "memory_in_mb": 256,
      "disk_in_mb": 1024,
      "health_check": {
        "type": "port",
        "data": {
          "timeout": null
        }
      },
      "created_at": "2018-05-22T21:06:42Z",
      "updated_at": "2018-06-05T21:07:08Z",
      "links": {
        "self": {
          "href": "https://api.run.example.com/v3/processes/30200dd1-64c0-451a-8c37-3c91e05c6741"
        },
        "scale": {
          "href": "https://api.run.example.com/v3/processes/30200dd1-64c0-451a-8c37-3c91e05c6741/actions/scale",
          "method": "POST"
        },
        "app": {
          "href": "https://api.run.example.com/v3/apps/30200dd1-64c0-451a-8c37-3c91e05c6741"
        },
        "space": {
          "href": "https://api.run.example.com/v3/spaces/60fa9e9b-c6eb-47a7-92ad-f438e775d234"
        },
        "stats": {
          "href": "https://api.run.example.com/v3/processes/30200dd1-64c0-451a-8c37-3c91e05c6741/stats"
        }
      }
    },
    {
      "guid": "0a8eb31c-0450-465d-90a1-837d92895504",
      "type": "web",
      "command": "[PRIVATE DATA HIDDEN IN LISTS]",
      "instances": 6,
      "memory_in_mb": 128,
      "disk_in_mb": 1024,
      "health_check": {
        "type": "port",
        "data": {
          "timeout": null
        }
      },
      "created_at": "2018-05-22T21:06:43Z",
      "updated_at": "2018-06-05T21:07:07Z",
      "links": {
        "self": {
          "href": "https://api.run.example.com/v3/processes/0a8eb31c-0450-465d-90a1-837d92895504"
        },
        "scale": {
          "href": "https://api.run.example.com/v3/processes/0a8eb31c-0450-465d-90a1-837d92895504/actions/scale",
          "method": "POST"
        },
        "app": {
          "href": "https://api.run.example.com/v3/apps/0a8eb31c-0450-465d-90a1-837d92895504"
        },
        "space": {
          "href": "https://api.run.example.com/v3/spaces/60fa9e9b-c6eb-47a7-92ad-f438e775d234"
        },
        "stats": {
          "href": "https://api.run.example.com/v3/processes/0a8eb31c-0450-465d-90a1-837d92895504/stats"
        }
      }
    },
    {
      "guid": "c55dca62-ff2e-4641-a4ec-9bea56f987b2",
      "type": "web",
      "command": "[PRIVATE DATA HIDDEN IN LISTS]",
      "instances": 1,
      "memory_in_mb": 1024,
      "disk_in_mb": 1024,
      "health_check": {
        "type": "port",
        "data": {
          "timeout": null
        }
      },
      "created_at": "2018-05-22T21:17:50Z",
      "updated_at": "2018-06-05T21:19:26Z",
      "links": {
        "self": {
          "href": "https://api.run.example.com/v3/processes/c55dca62-ff2e-4641-a4ec-9bea56f987b2"
        },
        "scale": {
          "href": "https://api.run.example.com/v3/processes/c55dca62-ff2e-4641-a4ec-9bea56f987b2/actions/scale",
          "method": "POST"
        },
        "app": {
          "href": "https://api.run.example.com/v3/apps/c55dca62-ff2e-4641-a4ec-9bea56f987b2"
        },
        "space": {
          "href": "https://api.run.example.com/v3/spaces/d5ecae4a-3bfc-4302-af74-bee452108316"
        },
        "stats": {
          "href": "https://api.run.example.com/v3/processes/c55dca62-ff2e-4641-a4ec-9bea56f987b2/stats"
        }
      }
    },
    {
      "guid": "72fab46a-8bcd-4941-b556-99694627c402",
      "type": "web",
      "command": "[PRIVATE DATA HIDDEN IN LISTS]",
      "instances": 1,
      "memory_in_mb": 1024,
      "disk_in_mb": 1024,
      "health_check": {
        "type": "port",
        "data": {
          "timeout": null
        }
      },
      "created_at": "2018-05-29T20:03:48Z",
      "updated_at": "2018-05-29T20:07:40Z",
      "links": {
        "self": {
          "href": "https://api.run.example.com/v3/processes/72fab46a-8bcd-4941-b556-99694627c402"
        },
        "scale": {
          "href": "https://api.run.example.com/v3/processes/72fab46a-8bcd-4941-b556-99694627c402/actions/scale",
          "method": "POST"
        },
        "app": {
          "href": "https://api.run.example.com/v3/apps/72fab46a-8bcd-4941-b556-99694627c402"
        },
        "space": {
          "href": "https://api.run.example.com/v3/spaces/6d6eaa00-1fba-4c31-9121-db9a72299240"
        },
        "stats": {
          "href": "https://api.run.example.com/v3/processes/72fab46a-8bcd-4941-b556-99694627c402/stats"
        }
      }
    },
    {
      "guid": "46b49bad-e469-418d-9f22-2fa8eb8a9144",
      "type": "web",
      "command": "[PRIVATE DATA HIDDEN IN LISTS]",
      "instances": 2,
      "memory_in_mb": 256,
      "disk_in_mb": 1024,
      "health_check": {
        "type": "port",
        "data": {
          "timeout": null
        }
      },
      "created_at": "2018-05-30T01:19:06Z",
      "updated_at": "2018-06-05T21:07:07Z",
      "links": {
        "self": {
          "href": "https://api.run.example.com/v3/processes/46b49bad-e469-418d-9f22-2fa8eb8a9144"
        },
        "scale": {
          "href": "https://api.run.example.com/v3/processes/46b49bad-e469-418d-9f22-2fa8eb8a9144/actions/scale",
          "method": "POST"
        },
        "app": {
          "href": "https://api.run.example.com/v3/apps/46b49bad-e469-418d-9f22-2fa8eb8a9144"
        },
        "space": {
          "href": "https://api.run.example.com/v3/spaces/60fa9e9b-c6eb-47a7-92ad-f438e775d234"
        },
        "stats": {
          "href": "https://api.run.example.com/v3/processes/46b49bad-e469-418d-9f22-2fa8eb8a9144/stats"
        }
      }
    },
    {
      "guid": "74600cfb-9cc9-4026-8473-f638f44ade12",
      "type": "web",
      "command": "[PRIVATE DATA HIDDEN IN LISTS]",
      "instances": 6,
      "memory_in_mb": 128,
      "disk_in_mb": 1024,
      "health_check": {
        "type": "port",
        "data": {
          "timeout": null
        }
      },
      "created_at": "2018-05-30T01:19:06Z",
      "updated_at": "2018-06-05T21:07:07Z",
      "links": {
        "self": {
          "href": "https://api.run.example.com/v3/processes/74600cfb-9cc9-4026-8473-f638f44ade12"
        },
        "scale": {
          "href": "https://api.run.example.com/v3/processes/74600cfb-9cc9-4026-8473-f638f44ade12/actions/scale",
          "method": "POST"
        },
        "app": {
          "href": "https://api.run.example.com/v3/apps/74600cfb-9cc9-4026-8473-f638f44ade12"
        },
        "space": {
          "href": "https://api.run.example.com/v3/spaces/60fa9e9b-c6eb-47a7-92ad-f438e775d234"
        },
        "stats": {
          "href": "https://api.run.example.com/v3/processes/74600cfb-9cc9-4026-8473-f638f44ade12/stats"
        }
      }
    },
    {
      "guid": "d782ee5f-8cff-451a-ad30-a0101c7fb430",
      "type": "web",
      "command": "[PRIVATE DATA HIDDEN IN LISTS]",
      "instances": 1,
      "memory_in_mb": 256,
      "disk_in_mb": 1024,
      "health_check": {
        "type": "port",
        "data": {
          "timeout": null
        }
      },
      "created_at": "2018-05-30T17:57:29Z",
      "updated_at": "2018-05-30T18:20:20Z",
      "links": {
        "self": {
          "href": "https://api.run.example.com/v3/processes/d782ee5f-8cff-451a-ad30-a0101c7fb430"
        },
        "scale": {
          "href": "https://api.run.example.com/v3/processes/d782ee5f-8cff-451a-ad30-a0101c7fb430/actions/scale",
          "method": "POST"
        },
        "app": {
          "href": "https://api.run.example.com/v3/apps/d782ee5f-8cff-451a-ad30-a0101c7fb430"
        },
        "space": {
          "href": "https://api.run.example.com/v3/spaces/6d6eaa00-1fba-4c31-9121-db9a72299240"
        },
        "stats": {
          "href": "https://api.run.example.com/v3/processes/d782ee5f-8cff-451a-ad30-a0101c7fb430/stats"
        }
      }
    },
    {
      "guid": "1b4d1907-3a54-46e6-ac73-d4732af2a28b",
      "type": "web",
      "command": "[PRIVATE DATA HIDDEN IN LISTS]",
      "instances": 2,
      "memory_in_mb": 1024,
      "disk_in_mb": 1024,
      "health_check": {
        "type": "port",
        "data": {
          "timeout": 360
        }
      },
      "created_at": "2018-06-05T19:41:09Z",
      "updated_at": "2018-06-05T21:06:35Z",
      "links": {
        "self": {
          "href": "https://api.run.example.com/v3/processes/1b4d1907-3a54-46e6-ac73-d4732af2a28b"
        },
        "scale": {
          "href": "https://api.run.example.com/v3/processes/1b4d1907-3a54-46e6-ac73-d4732af2a28b/actions/scale",
          "method": "POST"
        },
        "app": {
          "href": "https://api.run.example.com/v3/apps/1b4d1907-3a54-46e6-ac73-d4732af2a28b"
        },
        "space": {
          "href": "https://api.run.example.com/v3/spaces/60fa9e9b-c6eb-47a7-92ad-f438e775d234"
        },
        "stats": {
          "href": "https://api.run.example.com/v3/processes/1b4d1907-3a54-46e6-ac73-d4732af2a28b/stats"
        }
      }
    },
    {
      "guid": "bec57af6-4b84-4698-a366-1c4de1f4cc1e",
      "type": "web",
      "command": "[PRIVATE DATA HIDDEN IN LISTS]",
      "instances": 1,
      "memory_in_mb": 1024,
      "disk_in_mb": 1024,
      "health_check": {
        "type": "none",
        "data": {
          "timeout": null
        }
      },
      "created_at": "2018-06-05T19:41:09Z",
      "updated_at": "2018-06-05T21:06:14Z",
      "links": {
        "self": {
          "href": "https://api.run.example.com/v3/processes/bec57af6-4b84-4698-a366-1c4de1f4cc1e"
        },
        "scale": {
          "href": "https://api.run.example.com/v3/processes/bec57af6-4b84-4698-a366-1c4de1f4cc1e/actions/scale",
          "method": "POST"
        },
        "app": {
          "href": "https://api.run.example.com/v3/apps/bec57af6-4b84-4698-a366-1c4de1f4cc1e"
        },
        "space": {
          "href": "https://api.run.example.com/v3/spaces/60fa9e9b-c6eb-47a7-92ad-f438e775d234"
        },
        "stats": {
          "href": "https://api.run.example.com/v3/processes/bec57af6-4b84-4698-a366-1c4de1f4cc1e/stats"
        }
      }
    },
    {
      "guid": "5d449f8d-999a-4720-9296-5b6afce15a69",
      "type": "web",
      "command": "[PRIVATE DATA HIDDEN IN LISTS]",
      "instances": 1,
      "memory_in_mb": 2048,
      "disk_in_mb": 1024,
      "health_check": {
        "type": "none",
        "data": {
          "timeout": null
        }
      },
      "created_at": "2018-06-05T19:41:09Z",
      "updated_at": "2018-06-05T21:06:20Z",
      "links": {
        "self": {
          "href": "https://api.run.example.com/v3/processes/5d449f8d-999a-4720-9296-5b6afce15a69"
        },
        "scale": {
          "href": "https://api.run.example.com/v3/processes/5d449f8d-999a-4720-9296-5b6afce15a69/actions/scale",
          "method": "POST"
        },
        "app": {
          "href": "https://api.run.example.com/v3/apps/5d449f8d-999a-4720-9296-5b6afce15a69"
        },
        "space": {
          "href": "https://api.run.example.com/v3/spaces/60fa9e9b-c6eb-47a7-92ad-f438e775d234"
        },
        "stats": {
          "href": "https://api.run.example.com/v3/processes/5d449f8d-999a-4720-9296-5b6afce15a69/stats"
        }
      }
    },
    {
      "guid": "57a598af-c4a6-4ade-8db6-05679ae092f5",
      "type": "web",
      "command": "[PRIVATE DATA HIDDEN IN LISTS]",
      "instances": 2,
      "memory_in_mb": 1024,
      "disk_in_mb": 1024,
      "health_check": {
        "type": "port",
        "data": {
          "timeout": 180
        }
      },
      "created_at": "2018-06-05T19:45:25Z",
      "updated_at": "2018-06-05T21:10:23Z",
      "links": {
        "self": {
          "href": "https://api.run.example.com/v3/processes/57a598af-c4a6-4ade-8db6-05679ae092f5"
        },
        "scale": {
          "href": "https://api.run.example.com/v3/processes/57a598af-c4a6-4ade-8db6-05679ae092f5/actions/scale",
          "method": "POST"
        },
        "app": {
          "href": "https://api.run.example.com/v3/apps/57a598af-c4a6-4ade-8db6-05679ae092f5"
        },
        "space": {
          "href": "https://api.run.example.com/v3/spaces/1c17814c-e195-4ea5-9be2-6d01d9bf6007"
        },
        "stats": {
          "href": "https://api.run.example.com/v3/processes/57a598af-c4a6-4ade-8db6-05679ae092f5/stats"
        }
      }
    },
    {
      "guid": "79616d63-1075-4c02-a058-574a01314206",
      "type": "web",
      "command": "[PRIVATE DATA HIDDEN IN LISTS]",
      "instances": 1,
      "memory_in_mb": 1024,
      "disk_in_mb": 1024,
      "health_check": {
        "type": "none",
        "data": {
          "timeout": null
        }
      },
      "created_at": "2018-06-05T21:04:37Z",
      "updated_at": "2018-06-05T21:04:48Z",
      "links": {
        "self": {
          "href": "https://api.run.example.com/v3/processes/79616d63-1075-4c02-a058-574a01314206"
        },
        "scale": {
          "href": "https://api.run.example.com/v3/processes/79616d63-1075-4c02-a058-574a01314206/actions/scale",
          "method": "POST"
        },
        "app": {
          "href": "https://api.run.example.com/v3/apps/79616d63-1075-4c02-a058-574a01314206"
        },
        "space": {
          "href": "https://api.run.example.com/v3/spaces/60fa9e9b-c6eb-47a7-92ad-f438e775d234"
        },
        "stats": {
          "href": "https://api.run.example.com/v3/processes/79616d63-1075-4c02-a058-574a01314206/stats"
        }
      }
    },
    {
      "guid": "a2298d27-57fe-41e9-ad1e-f540427a0949",
      "type": "web",
      "command": "[PRIVATE DATA HIDDEN IN LISTS]",
      "instances": 2,
      "memory_in_mb": 1024,
      "disk_in_mb": 1024,
      "health_check": {
        "type": "port",
        "data": {
          "timeout": 360
        }
      },
      "created_at": "2018-06-05T21:04:37Z",
      "updated_at": "2018-06-05T21:06:34Z",
      "links": {
        "self": {
          "href": "https://api.run.example.com/v3/processes/a2298d27-57fe-41e9-ad1e-f540427a0949"
        },
        "scale": {
          "href": "https://api.run.example.com/v3/processes/a2298d27-57fe-41e9-ad1e-f540427a0949/actions/scale",
          "method": "POST"
        },
        "app": {
          "href": "https://api.run.example.com/v3/apps/a2298d27-57fe-41e9-ad1e-f540427a0949"
        },
        "space": {
          "href": "https://api.run.example.com/v3/spaces/60fa9e9b-c6eb-47a7-92ad-f438e775d234"
        },
        "stats": {
          "href": "https://api.run.example.com/v3/processes/a2298d27-57fe-41e9-ad1e-f540427a0949/stats"
        }
      }
    },
    {
      "guid": "a98f0649-4e31-435a-bb61-a96364c366a4",
      "type": "web",
      "command": "[PRIVATE DATA HIDDEN IN LISTS]",
      "instances": 1,
      "memory_in_mb": 2048,
      "disk_in_mb": 1024,
      "health_check": {
        "type": "none",
        "data": {
          "timeout": null
        }
      },
      "created_at": "2018-06-05T21:04:38Z",
      "updated_at": "2018-06-05T21:04:49Z",
      "links": {
        "self": {
          "href": "https://api.run.example.com/v3/processes/a98f0649-4e31-435a-bb61-a96364c366a4"
        },
        "scale": {
          "href": "https://api.run.example.com/v3/processes/a98f0649-4e31-435a-bb61-a96364c366a4/actions/scale",
          "method": "POST"
        },
        "app": {
          "href": "https://api.run.example.com/v3/apps/a98f0649-4e31-435a-bb61-a96364c366a4"
        },
        "space": {
          "href": "https://api.run.example.com/v3/spaces/60fa9e9b-c6eb-47a7-92ad-f438e775d234"
        },
        "stats": {
          "href": "https://api.run.example.com/v3/processes/a98f0649-4e31-435a-bb61-a96364c366a4/stats"
        }
      }
    },
    {
      "guid": "6aa28ed8-d377-4225-b853-80001e64113b",
      "type": "web",
      "command": "[PRIVATE DATA HIDDEN IN LISTS]",
      "instances": 3,
      "memory_in_mb": 64,
      "disk_in_mb": 1024,
      "health_check": {
        "type": "port",
        "data": {
          "timeout": null
        }
      },
      "created_at": "2018-06-05T21:07:29Z",
      "updated_at": "2018-06-05T21:08:02Z",
      "links": {
        "self": {
          "href": "https://api.run.example.com/v3/processes/6aa28ed8-d377-4225-b853-80001e64113b"
        },
        "scale": {
          "href": "https://api.run.example.com/v3/processes/6aa28ed8-d377-4225-b853-80001e64113b/actions/scale",
          "method": "POST"
        },
        "app": {
          "href": "https://api.run.example.com/v3/apps/6aa28ed8-d377-4225-b853-80001e64113b"
        },
        "space": {
          "href": "https://api.run.example.com/v3/spaces/2bd2080e-b9d7-4911-a607-a4e5adb00b59"
        },
        "stats": {
          "href": "https://api.run.example.com/v3/processes/6aa28ed8-d377-4225-b853-80001e64113b/stats"
        }
      }
    },
    {
      "guid": "dfc0f648-a3c8-4f46-b0a4-425136f238c6",
      "type": "web",
      "command": "[PRIVATE DATA HIDDEN IN LISTS]",
      "instances": 2,
      "memory_in_mb": 64,
      "disk_in_mb": 1024,
      "health_check": {
        "type": "port",
        "data": {
          "timeout": null
        }
      },
      "created_at": "2018-06-05T21:08:19Z",
      "updated_at": "2018-06-05T21:08:41Z",
      "links": {
        "self": {
          "href": "https://api.run.example.com/v3/processes/dfc0f648-a3c8-4f46-b0a4-425136f238c6"
        },
        "scale": {
          "href": "https://api.run.example.com/v3/processes/dfc0f648-a3c8-4f46-b0a4-425136f238c6/actions/scale",
          "method": "POST"
        },
        "app": {
          "href": "https://api.run.example.com/v3/apps/dfc0f648-a3c8-4f46-b0a4-425136f238c6"
        },
        "space": {
          "href": "https://api.run.example.com/v3/spaces/2bd2080e-b9d7-4911-a607-a4e5adb00b59"
        },
        "stats": {
          "href": "https://api.run.example.com/v3/processes/dfc0f648-a3c8-4f46-b0a4-425136f238c6/stats"
        }
      }
    },
    {
      "guid": "ba159a14-a3b5-4711-95c7-4bf89ff89a6c",
      "type": "web",
      "command": "[PRIVATE DATA HIDDEN IN LISTS]",
      "instances": 2,
      "memory_in_mb": 1024,
      "disk_in_mb": 1024,
      "health_check": {
        "type": "port",
        "data": {
          "timeout": 180
        }
      },
      "created_at": "2018-06-05T21:08:57Z",
      "updated_at": "2018-06-05T21:10:23Z",
      "links": {
        "self": {
          "href": "https://api.run.example.com/v3/processes/ba159a14-a3b5-4711-95c7-4bf89ff89a6c"
        },
        "scale": {
          "href": "https://api.run.example.com/v3/processes/ba159a14-a3b5-4711-95c7-4bf89ff89a6c/actions/scale",
          "method": "POST"
        },
        "app": {
          "href": "https://api.run.example.com/v3/apps/ba159a14-a3b5-4711-95c7-4bf89ff89a6c"
        },
        "space": {
          "href": "https://api.run.example.com/v3/spaces/1c17814c-e195-4ea5-9be2-6d01d9bf6007"
        },
        "stats": {
          "href": "https://api.run.example.com/v3/processes/ba159a14-a3b5-4711-95c7-4bf89ff89a6c/stats"
        }
      }
    },
    {
      "guid": "c7d20ea3-028b-4559-9f0c-a61966abc163",
      "type": "web",
      "command": "[PRIVATE DATA HIDDEN IN LISTS]",
      "instances": 3,
      "memory_in_mb": 256,
      "disk_in_mb": 1024,
      "health_check": {
        "type": "port",
        "data": {
          "timeout": null
        }
      },
      "created_at": "2018-06-05T21:11:44Z",
      "updated_at": "2018-06-05T21:12:00Z",
      "links": {
        "self": {
          "href": "https://api.run.example.com/v3/processes/c7d20ea3-028b-4559-9f0c-a61966abc163"
        },
        "scale": {
          "href": "https://api.run.example.com/v3/processes/c7d20ea3-028b-4559-9f0c-a61966abc163/actions/scale",
          "method": "POST"
        },
        "app": {
          "href": "https://api.run.example.com/v3/apps/c7d20ea3-028b-4559-9f0c-a61966abc163"
        },
        "space": {
          "href": "https://api.run.example.com/v3/spaces/7443e429-eee3-48ea-89cb-c9d3234c7fbf"
        },
        "stats": {
          "href": "https://api.run.example.com/v3/processes/c7d20ea3-028b-4559-9f0c-a61966abc163/stats"
        }
      }
    },
    {
      "guid": "956ca6e5-6560-49c6-ae01-843144b24ec5",
      "type": "web",
      "command": "[PRIVATE DATA HIDDEN IN LISTS]",
      "instances": 1,
      "memory_in_mb": 1024,
      "disk_in_mb": 1024,
      "health_check": {
        "type": "port",
        "data": {
          "timeout": 120
        }
      },
      "created_at": "2018-06-05T21:12:01Z",
      "updated_at": "2018-06-05T21:12:54Z",
      "links": {
        "self": {
          "href": "https://api.run.example.com/v3/processes/956ca6e5-6560-49c6-ae01-843144b24ec5"
        },
        "scale": {
          "href": "https://api.run.example.com/v3/processes/956ca6e5-6560-49c6-ae01-843144b24ec5/actions/scale",
          "method": "POST"
        },
        "app": {
          "href": "https://api.run.example.com/v3/apps/956ca6e5-6560-49c6-ae01-843144b24ec5"
        },
        "space": {
          "href": "https://api.run.example.com/v3/spaces/7443e429-eee3-48ea-89cb-c9d3234c7fbf"
        },
        "stats": {
          "href": "https://api.run.example.com/v3/processes/956ca6e5-6560-49c6-ae01-843144b24ec5/stats"
        }
      }
    },
    {
      "guid": "787fd0c1-d827-4c9e-975b-acd73267a583",
      "type": "web",
      "command": "[PRIVATE DATA HIDDEN IN LISTS]",
      "instances": 1,
      "memory_in_mb": 1024,
      "disk_in_mb": 1024,
      "health_check": {
        "type": "port",
        "data": {
          "timeout": null
        }
      },
      "created_at": "2018-06-29T18:31:10Z",
      "updated_at": "2018-06-29T18:31:10Z",
      "links": {
        "self": {
          "href": "https://api.run.example.com/v3/processes/787fd0c1-d827-4c9e-975b-acd73267a583"
        },
        "scale": {
          "href": "https://api.run.example.com/v3/processes/787fd0c1-d827-4c9e-975b-acd73267a583/actions/scale",
          "method": "POST"
        },
        "app": {
          "href": "https://api.run.example.com/v3/apps/787fd0c1-d827-4c9e-975b-acd73267a583"
        },
        "space": {
          "href": "https://api.run.example.com/v3/spaces/60fa9e9b-c6eb-47a7-92ad-f438e775d234"
        },
        "stats": {
          "href": "https://api.run.example.com/v3/processes/787fd0c1-d827-4c9e-975b-acd73267a583/stats"
        }
      }
    }
  ]
}`

const listProcessesPayload2 = `{
  "pagination": {
    "total_results": 26,
    "total_pages": 2,
    "first": {
      "href": "https://api.run.example.com/v3/processes?page=1&per_page=20"
    },
    "last": {
      "href": "https://api.run.example.com/v3/processes?page=2&per_page=20"
    },
    "next": null,
    "previous": {
      "href": "https://api.run.example.com/v3/processes?page=1&per_page=20"
    }
  },
  "resources": [
    {
      "guid": "09eb0d25-75b0-48e1-b45f-1636ea5bbe0c",
      "type": "web",
      "command": "[PRIVATE DATA HIDDEN IN LISTS]",
      "instances": 1,
      "memory_in_mb": 1024,
      "disk_in_mb": 1024,
      "health_check": {
        "type": "port",
        "data": {
          "timeout": null
        }
      },
      "created_at": "2018-06-29T18:32:35Z",
      "updated_at": "2018-06-29T18:32:35Z",
      "links": {
        "self": {
          "href": "https://api.run.example.com/v3/processes/09eb0d25-75b0-48e1-b45f-1636ea5bbe0c"
        },
        "scale": {
          "href": "https://api.run.example.com/v3/processes/09eb0d25-75b0-48e1-b45f-1636ea5bbe0c/actions/scale",
          "method": "POST"
        },
        "app": {
          "href": "https://api.run.example.com/v3/apps/09eb0d25-75b0-48e1-b45f-1636ea5bbe0c"
        },
        "space": {
          "href": "https://api.run.example.com/v3/spaces/60fa9e9b-c6eb-47a7-92ad-f438e775d234"
        },
        "stats": {
          "href": "https://api.run.example.com/v3/processes/09eb0d25-75b0-48e1-b45f-1636ea5bbe0c/stats"
        }
      }
    },
    {
      "guid": "00944be5-fe20-43ba-abe2-8fcc41ee2edc",
      "type": "web",
      "command": "[PRIVATE DATA HIDDEN IN LISTS]",
      "instances": 1,
      "memory_in_mb": 1024,
      "disk_in_mb": 1024,
      "health_check": {
        "type": "port",
        "data": {
          "timeout": null
        }
      },
      "created_at": "2018-07-13T20:17:23Z",
      "updated_at": "2018-07-18T17:20:59Z",
      "links": {
        "self": {
          "href": "https://api.run.example.com/v3/processes/00944be5-fe20-43ba-abe2-8fcc41ee2edc"
        },
        "scale": {
          "href": "https://api.run.example.com/v3/processes/00944be5-fe20-43ba-abe2-8fcc41ee2edc/actions/scale",
          "method": "POST"
        },
        "app": {
          "href": "https://api.run.example.com/v3/apps/00944be5-fe20-43ba-abe2-8fcc41ee2edc"
        },
        "space": {
          "href": "https://api.run.example.com/v3/spaces/a40da3ef-f0a3-4374-a7af-9c867382b30c"
        },
        "stats": {
          "href": "https://api.run.example.com/v3/processes/00944be5-fe20-43ba-abe2-8fcc41ee2edc/stats"
        }
      }
    },
    {
      "guid": "750ce4f3-3c28-4acc-9b50-357dc7dde27a",
      "type": "web",
      "command": "[PRIVATE DATA HIDDEN IN LISTS]",
      "instances": 1,
      "memory_in_mb": 1024,
      "disk_in_mb": 1024,
      "health_check": {
        "type": "port",
        "data": {
          "timeout": null
        }
      },
      "created_at": "2018-07-17T14:46:37Z",
      "updated_at": "2018-07-17T14:46:37Z",
      "links": {
        "self": {
          "href": "https://api.run.example.com/v3/processes/750ce4f3-3c28-4acc-9b50-357dc7dde27a"
        },
        "scale": {
          "href": "https://api.run.example.com/v3/processes/750ce4f3-3c28-4acc-9b50-357dc7dde27a/actions/scale",
          "method": "POST"
        },
        "app": {
          "href": "https://api.run.example.com/v3/apps/750ce4f3-3c28-4acc-9b50-357dc7dde27a"
        },
        "space": {
          "href": "https://api.run.example.com/v3/spaces/a40da3ef-f0a3-4374-a7af-9c867382b30c"
        },
        "stats": {
          "href": "https://api.run.example.com/v3/processes/750ce4f3-3c28-4acc-9b50-357dc7dde27a/stats"
        }
      }
    },
    {
      "guid": "2274c4c4-7e6e-413f-b976-e8210ddcc748",
      "type": "web",
      "command": "[PRIVATE DATA HIDDEN IN LISTS]",
      "instances": 1,
      "memory_in_mb": 256,
      "disk_in_mb": 1024,
      "health_check": {
        "type": "port",
        "data": {
          "timeout": null
        }
      },
      "created_at": "2018-08-06T21:45:14Z",
      "updated_at": "2018-08-07T18:11:33Z",
      "links": {
        "self": {
          "href": "https://api.run.example.com/v3/processes/2274c4c4-7e6e-413f-b976-e8210ddcc748"
        },
        "scale": {
          "href": "https://api.run.example.com/v3/processes/2274c4c4-7e6e-413f-b976-e8210ddcc748/actions/scale",
          "method": "POST"
        },
        "app": {
          "href": "https://api.run.example.com/v3/apps/2274c4c4-7e6e-413f-b976-e8210ddcc748"
        },
        "space": {
          "href": "https://api.run.example.com/v3/spaces/939eb497-671b-46cb-abd2-26e89129cc5b"
        },
        "stats": {
          "href": "https://api.run.example.com/v3/processes/2274c4c4-7e6e-413f-b976-e8210ddcc748/stats"
        }
      }
    },
    {
      "guid": "bc69b3fe-0d80-4508-a1d2-9bc53269287f",
      "type": "web",
      "command": "[PRIVATE DATA HIDDEN IN LISTS]",
      "instances": 1,
      "memory_in_mb": 128,
      "disk_in_mb": 1024,
      "health_check": {
        "type": "port",
        "data": {
          "timeout": null
        }
      },
      "created_at": "2018-08-07T22:27:14Z",
      "updated_at": "2018-08-08T16:09:41Z",
      "links": {
        "self": {
          "href": "https://api.run.example.com/v3/processes/bc69b3fe-0d80-4508-a1d2-9bc53269287f"
        },
        "scale": {
          "href": "https://api.run.example.com/v3/processes/bc69b3fe-0d80-4508-a1d2-9bc53269287f/actions/scale",
          "method": "POST"
        },
        "app": {
          "href": "https://api.run.example.com/v3/apps/bc69b3fe-0d80-4508-a1d2-9bc53269287f"
        },
        "space": {
          "href": "https://api.run.example.com/v3/spaces/a40da3ef-f0a3-4374-a7af-9c867382b30c"
        },
        "stats": {
          "href": "https://api.run.example.com/v3/processes/bc69b3fe-0d80-4508-a1d2-9bc53269287f/stats"
        }
      }
    },
    {
      "guid": "4b32fba5-ec34-40e1-9511-0a25c8f93d8f",
      "type": "worker",
      "command": "[PRIVATE DATA HIDDEN IN LISTS]",
      "instances": 1,
      "memory_in_mb": 1024,
      "disk_in_mb": 1024,
      "health_check": {
        "type": "process",
        "data": {
          "timeout": null
        }
      },
      "created_at": "2018-08-07T22:28:44Z",
      "updated_at": "2018-08-08T16:09:41Z",
      "links": {
        "self": {
          "href": "https://api.run.example.com/v3/processes/4b32fba5-ec34-40e1-9511-0a25c8f93d8f"
        },
        "scale": {
          "href": "https://api.run.example.com/v3/processes/4b32fba5-ec34-40e1-9511-0a25c8f93d8f/actions/scale",
          "method": "POST"
        },
        "app": {
          "href": "https://api.run.example.com/v3/apps/bc69b3fe-0d80-4508-a1d2-9bc53269287f"
        },
        "space": {
          "href": "https://api.run.example.com/v3/spaces/a40da3ef-f0a3-4374-a7af-9c867382b30c"
        },
        "stats": {
          "href": "https://api.run.example.com/v3/processes/4b32fba5-ec34-40e1-9511-0a25c8f93d8f/stats"
        }
      }
    }
  ]
}`

const createIsolationSegmentPayload = `{
   "guid": "323f211e-fea3-4161-9bd1-615392327913",
   "name": "TheKittenIsTheShark",
   "created_at": "2016-10-19T20:25:04Z",
   "updated_at": "2016-11-08T16:41:26Z",
   "links": {
      "self": {
         "href": "https://api.example.org/v3/isolation_segments/323f211e-fea3-4161-9bd1-615392327913"
      },
      "spaces": {
         "href": "https://api.example.org/v3/isolation_segments/323f211e-fea3-4161-9bd1-615392327913/relationships/spaces"
      },
      "organizations": {
         "href": "https://api.example.org/v3/isolation_segments/323f211e-fea3-4161-9bd1-615392327913/relationships/organizations"
      }
   }
}`

const listIsolationSegmentsPayload = `{
  "pagination": {
     "total_results": 2,
     "total_pages": 1,
     "first": {
        "href": "https://api.example.org/v3/isolation_segments?page=1&per_page=50"
     },
     "last": {
        "href": "https://api.example.org/v3/isolation_segments?page=1&per_page=50"
     },
     "next": null,
     "previous": null
  },
  "resources": [
     {
        "guid": "033b4c58-12bb-499a-b05d-4b6fc9e2993b",
        "name": "shared",
        "created_at": "2017-04-02T11:22:04Z",
        "updated_at": "2017-04-02T11:22:04Z",
        "links": {
           "self": {
              "href": "https://api.example.org/v3/isolation_segments/033b4c58-12bb-499a-b05d-4b6fc9e2993b"
           },
           "organizations": {
              "href": "https://api.example.org/v3/isolation_segments/033b4c58-12bb-499a-b05d-4b6fc9e2993b/organizations"
           },
           "spaces": {
              "href": "https://api.example.org/v3/isolation_segments/033b4c58-12bb-499a-b05d-4b6fc9e2993b/relationships/spaces"
           }
        }
     },
     {
        "guid": "23d0baf4-9d3c-44d8-b2dc-1767bcdad1e0",
        "name": "my_segment",
        "created_at": "2017-04-07T11:20:16Z",
        "updated_at": "2017-04-07T11:20:16Z",
        "links": {
           "self": {
              "href": "https://api.example.org/v3/isolation_segments/23d0baf4-9d3c-44d8-b2dc-1767bcdad1e0"
           },
           "organizations": {
              "href": "https://api.example.org/v3/isolation_segments/23d0baf4-9d3c-44d8-b2dc-1767bcdad1e0/organizations"
           },
           "spaces": {
              "href": "https://api.example.org/v3/isolation_segments/23d0baf4-9d3c-44d8-b2dc-1767bcdad1e0/relationships/spaces"
           }
        }
     }
  ]
}`

const listIsolationSegmentsPayloadPage1 = `{
   "pagination": {
      "total_results": 4,
      "total_pages": 2,
      "first": {
         "href": "https://api.example.org/v3/isolation_segments?page=1&per_page=2"
      },
      "last": {
         "href": "https://api.example.org/v3/isolation_segments?page=2&per_page=2"
      },
      "next": {
        "href": "https://api.example.org/v3/isolation_segments?page=2&per_page=2"
      },
      "previous": null
   },
   "resources": [
      {
         "guid": "033b4c58-12bb-499a-b05d-4b6fc9e2993b",
         "name": "shared",
         "created_at": "2017-04-02T11:22:04Z",
         "updated_at": "2017-04-02T11:22:04Z",
         "links": {
            "self": {
               "href": "https://api.example.org/v3/isolation_segments/033b4c58-12bb-499a-b05d-4b6fc9e2993b"
            },
            "organizations": {
               "href": "https://api.example.org/v3/isolation_segments/033b4c58-12bb-499a-b05d-4b6fc9e2993b/organizations"
            },
            "spaces": {
               "href": "https://api.example.org/v3/isolation_segments/033b4c58-12bb-499a-b05d-4b6fc9e2993b/relationships/spaces"
            }
         }
      },
      {
         "guid": "23d0baf4-9d3c-44d8-b2dc-1767bcdad1e0",
         "name": "my_segment",
         "created_at": "2017-04-07T11:20:16Z",
         "updated_at": "2017-04-07T11:20:16Z",
         "links": {
            "self": {
               "href": "https://api.example.org/v3/isolation_segments/23d0baf4-9d3c-44d8-b2dc-1767bcdad1e0"
            },
            "organizations": {
               "href": "https://api.example.org/v3/isolation_segments/23d0baf4-9d3c-44d8-b2dc-1767bcdad1e0/organizations"
            },
            "spaces": {
               "href": "https://api.example.org/v3/isolation_segments/23d0baf4-9d3c-44d8-b2dc-1767bcdad1e0/relationships/spaces"
            }
         }
      }
   ]
}`

const listIsolationSegmentsPayloadPage2 = `{
  "pagination": {
     "total_results": 4,
     "total_pages": 2,
     "first": {
        "href": "https://api.example.org/v3/isolation_segments?page=1&per_page=2"
      },
      "last": {
        "href": "https://api.example.org/v3/isolation_segments?page=2&per_page=2"
      },
     "next": null,
     "previous": {
        "href": "https://api.example.org/v3/isolation_segments?page=1&per_page=2"
      }
  },
  "resources": [
     {
        "guid": "abcdefg12-12bb-499a-b05d-4b6fc9e2993b",
        "name": "shared1",
        "created_at": "2017-05-02T11:22:04Z",
        "updated_at": "2017-05-02T11:22:04Z",
        "links": {
           "self": {
              "href": "https://api.example.org/v3/isolation_segments/abcdefg12-12bb-499a-b05d-4b6fc9e2993b"
           },
           "organizations": {
              "href": "https://api.example.org/v3/isolation_segments/abcdefg12-12bb-499a-b05d-4b6fc9e2993b/organizations"
           },
           "spaces": {
              "href": "https://api.example.org/v3/isolation_segments/abcdefg12-12bb-499a-b05d-4b6fc9e2993b/relationships/spaces"
           }
        }
     },
     {
        "guid": "abcdef123-9d3c-44d8-b2dc-1767bcdad1e0",
        "name": "my_segment1",
        "created_at": "2017-05-07T11:20:16Z",
        "updated_at": "2017-05-07T11:20:16Z",
        "links": {
           "self": {
              "href": "https://api.example.org/v3/isolation_segments/abcdef123-9d3c-44d8-b2dc-1767bcdad1e0"
           },
           "organizations": {
              "href": "https://api.example.org/v3/isolation_segments/abcdef123-9d3c-44d8-b2dc-1767bcdad1e0/organizations"
           },
           "spaces": {
              "href": "https://api.example.org/v3/isolation_segments/abcdef123-9d3c-44d8-b2dc-1767bcdad1e0/relationships/spaces"
           }
        }
     }
  ]
}`

const listV3SpaceRolesBySpaceGUIDPayload = `{
  "pagination": {
    "total_results": 3,
    "total_pages": 2,
    "first": {
      "href": "https://api.example.org/v3/roles?page=1&per_page=2&space_guids=spaceGUID1"
    },
    "last": {
      "href": "https://api.example.org/v3/roles?page=2&per_page=2&space_guids=spaceGUID1"
    },
    "next": {
      "href": "https://api.example.org/v3/rolespage2?page=2&per_page=2&space_guids=spaceGUID1"
    },
    "previous": null
  },
  "resources": [
    {
      "guid": "roleGUID1",
      "created_at": "2019-10-10T17:19:12Z",
      "updated_at": "2019-10-10T17:19:12Z",
      "type": "space_developer",
      "relationships": {
        "user": {
          "data": {
            "guid": "userGUID1"
          }
        },
        "space": {
          "data": {
            "guid": "spaceGUID1"
          }
        },
        "organization": {
          "data": null
        }
       },
       "links": {
          "self": {
            "href": "https://api.example.org/v3/roles/roleGUID1"
          },
          "user": {
            "href": "https://api.example.org/v3/users/userGUID1"
          },
          "space": {
            "href": "https://api.example.org/v3/spaces/spaceGUID1"
          }
       }
    },
    {
      "guid": "roleGUID2",
      "created_at": "2047-11-10T17:19:12Z",
      "updated_at": "2047-11-10T17:19:12Z",
      "type": "space_auditor",
      "relationships": {
        "user": {
          "data": {
            "guid": "userGUID2"
          }
        },
        "space": {
          "data": {
            "guid": "spaceGUID1"
          }
        },
        "organization": {
          "data": null
        }
      },
      "links": {
        "self": {
          "href": "https://api.example.org/v3/roles/roleGUID2"
        },
        "user": {
          "href": "https://api.example.org/v3/users/userGUID2"
        },
        "space": {
          "href": "https://api.example.org/v3/spaces/spaceGUID1"
        }
      }
    }
  ]
}`

const listV3SpaceRolesBySpaceGuidPayloadPage2 = `{
  "pagination": {
    "total_results": 3,
    "total_pages": 2,
    "first": {
      "href": "https://api.example.org/v3/roles?page=1&per_page=2&space_guids=spaceGUID1"
    },
    "last": {
      "href": "https://api.example.org/v3/rolespage2?page=2&per_page=2&space_guids=spaceGUID1"
    },
    "next": null,
    "previous": {
      "href": "https://api.example.org/v3/roles?page=1&per_page=2&space_guids=spaceGUID1"
    }
  },
  "resources": [
    {
      "guid": "roleGUID3",
      "created_at": "2019-10-10T17:19:12Z",
      "updated_at": "2019-10-10T17:19:12Z",
      "type": "space_manager",
      "relationships": {
        "user": {
          "data": {
            "guid": "userGUID2"
          }
        },
        "space": {
          "data": {
            "guid": "spaceGUID1"
          }
        },
        "organization": {
          "data": null
        }
      },
      "links": {
        "self": {
          "href": "https://api.example.org/v3/roles/roleGUID3"
        },
        "user": {
          "href": "https://api.example.org/v3/users/userGUID2"
        },
        "space": {
          "href": "https://api.example.org/v3/spaces/spaceGUID1"
        }
      }
    }
  ]
}`

const listV3SpaceRoleUsersBySpaceGUIDPayload = `{
  "pagination": {
    "total_results": 3,
    "total_pages": 2,
    "first": {
      "href": "https://api.example.org/v3/roles?page=1&per_page=2&include=user&space_guids=spaceGUID1"
    },
    "last": {
      "href": "https://api.example.org/v3/roles?page=2&per_page=2&include=user&space_guids=spaceGUID1"
    },
    "next": {
      "href": "https://api.example.org/v3/rolespage2?page=2&per_page=2&include=user&space_guids=spaceGUID1"
    },
    "previous": null
  },
  "resources": [
    {
      "guid": "roleGUID1",
      "created_at": "2019-10-10T17:19:12Z",
      "updated_at": "2019-10-10T17:19:12Z",
      "type": "space_supporter",
      "relationships": {
        "user": {
          "data": {
            "guid": "userGUID1"
          }
        },
        "space": {
          "data": {
            "guid": "spaceGUID1"
          }
        },
        "organization": {
          "data": null
        }
       },
       "links": {
          "self": {
            "href": "https://api.example.org/v3/roles/roleGUID1"
          },
          "user": {
            "href": "https://api.example.org/v3/users/userGUID1"
          },
          "space": {
            "href": "https://api.example.org/v3/spaces/spaceGUID1"
          }
       }
    },
    {
      "guid": "roleGUID2",
      "created_at": "2047-11-10T17:19:12Z",
      "updated_at": "2047-11-10T17:19:12Z",
      "type": "space_supporter",
      "relationships": {
        "user": {
          "data": {
            "guid": "userGUID2"
          }
        },
        "space": {
          "data": {
            "guid": "spaceGUID1"
          }
        },
        "organization": {
          "data": null
        }
      },
      "links": {
        "self": {
          "href": "https://api.example.org/v3/roles/roleGUID2"
        },
        "user": {
          "href": "https://api.example.org/v3/users/userGUID2"
        },
        "space": {
          "href": "https://api.example.org/v3/spaces/spaceGUID1"
        }
      }
    }
  ],
  "included": {
    "users": [
      {
        "guid": "userGUID1",
        "created_at": "2022-05-25T23:57:45Z",
        "updated_at": "2022-05-25T23:57:45Z",
        "username": "user1",
        "presentation_name": "user1",
        "origin": "uaa",
        "metadata": {
            "labels": {},
            "annotations": {}
        },
        "links": {
            "self": {
              "href": "https://api.example.org/v3/users/userGUID1"
            }
        }
      },
      {
        "guid": "userGUID2",
        "created_at": "2022-05-25T23:57:45Z",
        "updated_at": "2022-05-25T23:57:45Z",
        "username": "user2",
        "presentation_name": "user2",
        "origin": "uaa",
        "metadata": {
            "labels": {},
            "annotations": {}
        },
        "links": {
            "self": {
              "href": "https://api.example.org/v3/users/userGUID2"
            }
        }
      }
    ]
  }     
}`

const listV3SpaceRoleUsersBySpaceGUIDPayloadPage2 = `{
  "pagination": {
    "total_results": 3,
    "total_pages": 2,
    "first": {
      "href": "https://api.example.org/v3/roles?page=1&per_page=2&include=user&space_guids=spaceGUID1"
    },
    "last": {
      "href": "https://api.example.org/v3/rolespage2?page=2&per_page=2&include=user&space_guids=spaceGUID1"
    },
    "next": null,
    "previous": {
      "href": "https://api.example.org/v3/roles?page=1&per_page=2&include=user&space_guids=spaceGUID1"
    }
  },
  "resources": [
    {
      "guid": "roleGUID3",
      "created_at": "2019-10-10T17:19:12Z",
      "updated_at": "2019-10-10T17:19:12Z",
      "type": "space_supporter",
      "relationships": {
        "user": {
          "data": {
            "guid": "userGUID2"
          }
        },
        "space": {
          "data": {
            "guid": "spaceGUID1"
          }
        },
        "organization": {
          "data": null
        }
      },
      "links": {
        "self": {
          "href": "https://api.example.org/v3/roles/roleGUID3"
        },
        "user": {
          "href": "https://api.example.org/v3/users/userGUID2"
        },
        "space": {
          "href": "https://api.example.org/v3/spaces/spaceGUID1"
        }
      }
    }
  ],
  "included": {
    "users": [
      {
        "guid": "userGUID3",
        "created_at": "2022-05-25T23:57:45Z",
        "updated_at": "2022-05-25T23:57:45Z",
        "username": "user3",
        "presentation_name": "user3",
        "origin": "uaa",
        "metadata": {
            "labels": {},
            "annotations": {}
        },
        "links": {
            "self": {
              "href": "https://api.example.org/v3/users/userGUID3"
            }
        }
      }
    ]
  } 
}`

const listV3SpaceRolesBySpaceGUIDAndTypePayload = `{
  "pagination": {
    "total_results": 3,
    "total_pages": 2,
    "first": {
      "href": "https://api.example.org/v3/roles?page=1&per_page=2&include=user&space_guids=spaceGUID1&types=space_supporter"
    },
    "last": {
      "href": "https://api.example.org/v3/roles?page=2&per_page=2&include=user&space_guids=spaceGUID1&types=space_supporter"
    },
    "next": {
      "href": "https://api.example.org/v3/rolespage2?page=2&per_page=2&include=user&space_guids=spaceGUID1&types=space_supporter"
    },
    "previous": null
  },
  "resources": [
    {
      "guid": "roleGUID1",
      "created_at": "2019-10-10T17:19:12Z",
      "updated_at": "2019-10-10T17:19:12Z",
      "type": "space_supporter",
      "relationships": {
        "user": {
          "data": {
            "guid": "userGUID1"
          }
        },
        "space": {
          "data": {
            "guid": "spaceGUID1"
          }
        },
        "organization": {
          "data": null
        }
       },
       "links": {
          "self": {
            "href": "https://api.example.org/v3/roles/roleGUID1"
          },
          "user": {
            "href": "https://api.example.org/v3/users/userGUID1"
          },
          "space": {
            "href": "https://api.example.org/v3/spaces/spaceGUID1"
          }
       }
    },
    {
      "guid": "roleGUID2",
      "created_at": "2047-11-10T17:19:12Z",
      "updated_at": "2047-11-10T17:19:12Z",
      "type": "space_supporter",
      "relationships": {
        "user": {
          "data": {
            "guid": "userGUID2"
          }
        },
        "space": {
          "data": {
            "guid": "spaceGUID1"
          }
        },
        "organization": {
          "data": null
        }
      },
      "links": {
        "self": {
          "href": "https://api.example.org/v3/roles/roleGUID2"
        },
        "user": {
          "href": "https://api.example.org/v3/users/userGUID2"
        },
        "space": {
          "href": "https://api.example.org/v3/spaces/spaceGUID1"
        }
      }
    }
  ],
  "included": {
    "users": [
      {
        "guid": "userGUID1",
        "created_at": "2022-05-25T23:57:45Z",
        "updated_at": "2022-05-25T23:57:45Z",
        "username": "user1",
        "presentation_name": "user1",
        "origin": "uaa",
        "metadata": {
            "labels": {},
            "annotations": {}
        },
        "links": {
            "self": {
              "href": "https://api.example.org/v3/users/userGUID1"
            }
        }
      },
      {
        "guid": "userGUID2",
        "created_at": "2022-05-25T23:57:45Z",
        "updated_at": "2022-05-25T23:57:45Z",
        "username": "user2",
        "presentation_name": "user2",
        "origin": "uaa",
        "metadata": {
            "labels": {},
            "annotations": {}
        },
        "links": {
            "self": {
              "href": "https://api.example.org/v3/users/userGUID2"
            }
        }
      }
    ]
  }     
}`

const listV3SpaceRolesBySpaceGuidAndTypePayloadPage2 = `{
  "pagination": {
    "total_results": 3,
    "total_pages": 2,
    "first": {
      "href": "https://api.example.org/v3/roles?page=1&per_page=2&include=user&space_guids=spaceGUID1&types=space_supporter"
    },
    "last": {
      "href": "https://api.example.org/v3/rolespage2?page=2&per_page=2&include=user&space_guids=spaceGUID1&types=space_supporter"
    },
    "next": null,
    "previous": {
      "href": "https://api.example.org/v3/roles?page=1&per_page=2&include=user&space_guids=spaceGUID1&types=space_supporter"
    }
  },
  "resources": [
    {
      "guid": "roleGUID3",
      "created_at": "2019-10-10T17:19:12Z",
      "updated_at": "2019-10-10T17:19:12Z",
      "type": "space_supporter",
      "relationships": {
        "user": {
          "data": {
            "guid": "userGUID2"
          }
        },
        "space": {
          "data": {
            "guid": "spaceGUID1"
          }
        },
        "organization": {
          "data": null
        }
      },
      "links": {
        "self": {
          "href": "https://api.example.org/v3/roles/roleGUID3"
        },
        "user": {
          "href": "https://api.example.org/v3/users/userGUID2"
        },
        "space": {
          "href": "https://api.example.org/v3/spaces/spaceGUID1"
        }
      }
    }
  ],
  "included": {
    "users": [
      {
        "guid": "userGUID3",
        "created_at": "2022-05-25T23:57:45Z",
        "updated_at": "2022-05-25T23:57:45Z",
        "username": "user3",
        "presentation_name": "user3",
        "origin": "uaa",
        "metadata": {
            "labels": {},
            "annotations": {}
        },
        "links": {
            "self": {
              "href": "https://api.example.org/v3/users/userGUID3"
            }
        }
      }
    ]
  } 
}`

const listV3OrganizationRolesByOrganizationGUIDPayload = `{
  "pagination": {
    "total_results": 3,
    "total_pages": 2,
    "first": {
      "href": "https://api.example.org/v3/roles?page=1&per_page=2&include=user&organization_guids=orgGUID1"
    },
    "last": {
      "href": "https://api.example.org/v3/roles?page=2&per_page=2&include=user&organization_guids=orgGUID1"
    },
    "next": {
      "href": "https://api.example.org/v3/rolespage2?page=2&per_page=2&include=user&organization_guids=orgGUID1"
    },
    "previous": null
  },
  "resources": [
    {
      "guid": "roleGUID1",
      "created_at": "2019-10-10T17:19:12Z",
      "updated_at": "2019-10-10T17:19:12Z",
      "type": "organization_auditor",
      "relationships": {
        "user": {
          "data": {
            "guid": "userGUID1"
          }
        },
        "space": {
          "data": null
        },
        "organization": {
          "data": {
            "guid": "orgGUID1"
          }
        }
       },
       "links": {
          "self": {
            "href": "https://api.example.org/v3/roles/roleGUID1"
          },
          "user": {
            "href": "https://api.example.org/v3/users/userGUID1"
          },
          "org": {
            "href": "https://api.example.org/v3/organization/orgGUID1"
          }
       }
    },
    {
      "guid": "roleGUID2",
      "created_at": "2047-11-10T17:19:12Z",
      "updated_at": "2047-11-10T17:19:12Z",
      "type": "organization_auditor",
      "relationships": {
        "user": {
          "data": {
            "guid": "userGUID2"
          }
        },
        "space": {
          "data": null
        },
        "organization": {
          "data": {
            "guid": "orgGUID1"
          }
        }
      },
      "links": {
        "self": {
          "href": "https://api.example.org/v3/roles/roleGUID2"
        },
        "user": {
          "href": "https://api.example.org/v3/users/userGUID2"
        },
        "organization": {
          "href": "https://api.example.org/v3/organization/orgGUID1"
        }
      }
    }
  ],
  "included": {
    "users": [
      {
        "guid": "userGUID1",
        "created_at": "2022-05-25T23:57:45Z",
        "updated_at": "2022-05-25T23:57:45Z",
        "username": "user1",
        "presentation_name": "user1",
        "origin": "uaa",
        "metadata": {
            "labels": {},
            "annotations": {}
        },
        "links": {
            "self": {
              "href": "https://api.example.org/v3/users/userGUID1"
            }
        }
      },
      {
        "guid": "userGUID2",
        "created_at": "2022-05-25T23:57:45Z",
        "updated_at": "2022-05-25T23:57:45Z",
        "username": "user2",
        "presentation_name": "user2",
        "origin": "uaa",
        "metadata": {
            "labels": {},
            "annotations": {}
        },
        "links": {
            "self": {
              "href": "https://api.example.org/v3/users/userGUID2"
            }
        }
      }
    ]
  }     
}`

const listV3OrganizationRolesByOrganizationGuidPayloadPage2 = `{
  "pagination": {
    "total_results": 3,
    "total_pages": 2,
    "first": {
      "href": "https://api.example.org/v3/roles?page=1&per_page=2&include=user&organiziation_guids=orgGUID1"
    },
    "last": {
      "href": "https://api.example.org/v3/rolespage2?page=2&per_page=2&include=user&organiziation_guids=orgGUID1"
    },
    "next": null,
    "previous": {
      "href": "https://api.example.org/v3/roles?page=1&per_page=2&include=user&organiziation_guids=orgGUID1"
    }
  },
  "resources": [
    {
      "guid": "roleGUID3",
      "created_at": "2019-10-10T17:19:12Z",
      "updated_at": "2019-10-10T17:19:12Z",
      "type": "organization_auditor",
      "relationships": {
        "user": {
          "data": {
            "guid": "userGUID2"
          }
        },
        "space": {
          "data": null
        },
        "organization": {
          "data": {
            "guid": "spaceGUID1"
          }
        }
      },
      "links": {
        "self": {
          "href": "https://api.example.org/v3/roles/roleGUID3"
        },
        "user": {
          "href": "https://api.example.org/v3/users/userGUID2"
        },
        "organization": {
          "href": "https://api.example.org/v3/organization/orgGUID1"
        }
      }
    }
  ],
  "included": {
    "users": [
      {
        "guid": "userGUID3",
        "created_at": "2022-05-25T23:57:45Z",
        "updated_at": "2022-05-25T23:57:45Z",
        "username": "user3",
        "presentation_name": "user3",
        "origin": "uaa",
        "metadata": {
            "labels": {},
            "annotations": {}
        },
        "links": {
            "self": {
              "href": "https://api.example.org/v3/users/userGUID3"
            }
        }
      }
    ]
  } 
}`

const listV3OrganizationRolesByOrganizationGUIDAndTypePayload = `{
  "pagination": {
    "total_results": 3,
    "total_pages": 2,
    "first": {
      "href": "https://api.example.org/v3/roles?page=1&per_page=2&include=user&organization_guids=orgGUID1&types=organization_auditor"
    },
    "last": {
      "href": "https://api.example.org/v3/roles?page=2&per_page=2&include=user&organization_guids=orgGUID1&types=organization_auditor"
    },
    "next": {
      "href": "https://api.example.org/v3/rolespage2?page=2&per_page=2&include=user&organization_guids=orgGUID1&types=organization_auditor"
    },
    "previous": null
  },
  "resources": [
    {
      "guid": "roleGUID1",
      "created_at": "2019-10-10T17:19:12Z",
      "updated_at": "2019-10-10T17:19:12Z",
      "type": "organization_auditor",
      "relationships": {
        "user": {
          "data": {
            "guid": "userGUID1"
          }
        },
        "space": {
          "data": null
        },
        "organization": {
          "data": {
            "guid": "orgGUID1"
          }
        }
       },
       "links": {
          "self": {
            "href": "https://api.example.org/v3/roles/roleGUID1"
          },
          "user": {
            "href": "https://api.example.org/v3/users/userGUID1"
          },
          "org": {
            "href": "https://api.example.org/v3/organization/orgGUID1"
          }
       }
    },
    {
      "guid": "roleGUID2",
      "created_at": "2047-11-10T17:19:12Z",
      "updated_at": "2047-11-10T17:19:12Z",
      "type": "organization_auditor",
      "relationships": {
        "user": {
          "data": {
            "guid": "userGUID2"
          }
        },
        "space": {
          "data": null
        },
        "organization": {
          "data": {
            "guid": "orgGUID1"
          }
        }
      },
      "links": {
        "self": {
          "href": "https://api.example.org/v3/roles/roleGUID2"
        },
        "user": {
          "href": "https://api.example.org/v3/users/userGUID2"
        },
        "organization": {
          "href": "https://api.example.org/v3/organization/orgGUID1"
        }
      }
    }
  ],
  "included": {
    "users": [
      {
        "guid": "userGUID1",
        "created_at": "2022-05-25T23:57:45Z",
        "updated_at": "2022-05-25T23:57:45Z",
        "username": "user1",
        "presentation_name": "user1",
        "origin": "uaa",
        "metadata": {
            "labels": {},
            "annotations": {}
        },
        "links": {
            "self": {
              "href": "https://api.example.org/v3/users/userGUID1"
            }
        }
      },
      {
        "guid": "userGUID2",
        "created_at": "2022-05-25T23:57:45Z",
        "updated_at": "2022-05-25T23:57:45Z",
        "username": "user2",
        "presentation_name": "user2",
        "origin": "uaa",
        "metadata": {
            "labels": {},
            "annotations": {}
        },
        "links": {
            "self": {
              "href": "https://api.example.org/v3/users/userGUID2"
            }
        }
      }
    ]
  }     
}`

const listV3OrganizationRolesByOrganizationGuidAndTypePayloadPage2 = `{
  "pagination": {
    "total_results": 3,
    "total_pages": 2,
    "first": {
      "href": "https://api.example.org/v3/roles?page=1&per_page=2&include=user&organiziation_guids=orgGUID1&types=organization_auditor"
    },
    "last": {
      "href": "https://api.example.org/v3/rolespage2?page=2&per_page=2&include=user&organiziation_guids=orgGUID1&types=organization_auditor"
    },
    "next": null,
    "previous": {
      "href": "https://api.example.org/v3/roles?page=1&per_page=2&include=user&organiziation_guids=orgGUID1&types=organization_auditor"
    }
  },
  "resources": [
    {
      "guid": "roleGUID3",
      "created_at": "2019-10-10T17:19:12Z",
      "updated_at": "2019-10-10T17:19:12Z",
      "type": "organization_auditor",
      "relationships": {
        "user": {
          "data": {
            "guid": "userGUID2"
          }
        },
        "space": {
          "data": null
        },
        "organization": {
          "data": {
            "guid": "spaceGUID1"
          }
        }
      },
      "links": {
        "self": {
          "href": "https://api.example.org/v3/roles/roleGUID3"
        },
        "user": {
          "href": "https://api.example.org/v3/users/userGUID2"
        },
        "organization": {
          "href": "https://api.example.org/v3/organization/orgGUID1"
        }
      }
    }
  ],
  "included": {
    "users": [
      {
        "guid": "userGUID3",
        "created_at": "2022-05-25T23:57:45Z",
        "updated_at": "2022-05-25T23:57:45Z",
        "username": "user3",
        "presentation_name": "user3",
        "origin": "uaa",
        "metadata": {
            "labels": {},
            "annotations": {}
        },
        "links": {
            "self": {
              "href": "https://api.example.org/v3/users/userGUID3"
            }
        }
      }
    ]
  } 
}`

const listV3SpaceRolesByUserGuidPayload = `{
  "pagination": {
    "total_results": 2,
    "total_pages": 1,
    "first": {
      "href": "https://api.example.org/v3/roles?page=1&per_page=2&user_guids=userGUID1"
    },
    "last": {
      "href": "https://api.example.org/v3/roles?page=2&per_page=2&user_guids=userGUID1"
    },
    "next": null,
    "previous": null
  },
  "resources": [
    {
      "guid": "roleGUID1",
      "created_at": "2019-10-10T17:19:12Z",
      "updated_at": "2019-10-10T17:19:12Z",
      "type": "space_developer",
      "relationships": {
        "user": {
          "data": {
            "guid": "userGUID1"
          }
        },
        "space": {
          "data": {
            "guid": "spaceGUID1"
          }
        },
        "organization": {
          "data": null
        }
       },
       "links": {
          "self": {
            "href": "https://api.example.org/v3/roles/roleGUID1"
          },
          "user": {
            "href": "https://api.example.org/v3/users/userGUID1"
          },
          "space": {
            "href": "https://api.example.org/v3/spaces/spaceGUID1"
          }
       }
    },
    {
      "guid": "roleGUID4",
      "created_at": "2019-10-10T17:19:12Z",
      "updated_at": "2019-10-10T17:19:12Z",
      "type": "space_manager",
      "relationships": {
        "user": {
          "data": {
            "guid": "userGUID1"
          }
        },
        "space": {
          "data": {
            "guid": "spaceGUID2"
          }
        },
        "organization": {
          "data": null
        }
      },
      "links": {
        "self": {
          "href": "https://api.example.org/v3/roles/roleGUID4"
        },
        "user": {
          "href": "https://api.example.org/v3/users/userGUID1"
        },
        "space": {
          "href": "https://api.example.org/v3/spaces/spaceGUID2"
        }
      }
    }
  ]
}`

const listV3spaceRolesBySpaceAndUserGuidPayload = `{
  "pagination": {
    "total_results": 1,
    "total_pages": 1,
    "first": {
      "href": "https://api.example.org/v3/roles?page=1&per_page=1&space_guids=spaceGUID2&user_guids=userGUID1"
    },
    "last": {
      "href": "https://api.example.org/v3/roles?page=1&per_page=1&space_guids=spaceGUID2&user_guids=userGUID1"
    },
    "next": null,
    "previous": null
  },
  "resources": [
    {
      "guid": "roleGUID4",
      "created_at": "2019-10-10T17:19:12Z",
      "updated_at": "2019-10-10T17:19:12Z",
      "type": "space_manager",
      "relationships": {
        "user": {
          "data": {
            "guid": "userGUID1"
          }
        },
        "space": {
          "data": {
            "guid": "spaceGUID2"
          }
        },
        "organization": {
          "data": null
        }
      },
      "links": {
        "self": {
          "href": "https://api.example.org/v3/roles/roleGUID4"
        },
        "user": {
          "href": "https://api.example.org/v3/users/userGUID1"
        },
        "space": {
          "href": "https://api.example.org/v3/spaces/spaceGUID2"
        }
      }
    }
  ]
}`

const listV3SecurityGroupsPayload = `{
  "pagination": {
    "total_results": 2,
    "total_pages": 1,
    "first": {
      "href": "https://api.example.org/v3/security_groups?page=1&per_page=50"
    },
    "last": {
      "href": "https://api.example.org/v3/security_groups?page=1&per_page=50"
    },
    "next": null,
    "previous": null
  },
  "resources": [
    {
      "guid": "guid-1",
      "name": "my-group1",
      "globally_enabled": {
        "running": true,
        "staging": false
      },
      "rules": [
        {
          "protocol": "tcp",
          "destination": "1.2.3.4/10",
          "ports": "443,80,8080"
        },
        {
          "protocol": "icmp",
          "destination": "1.2.3.4/12",
          "type": 8,
          "code": 0,
          "description": "test-desc-1"
        }
      ],
      "relationships": {
        "staging_spaces": {
          "data": [
            { "guid": "space-guid-1" },
            { "guid": "space-guid-2" }
          ]
        },
        "running_spaces": {
          "data": []
        }
      },
      "links": {
        "self": {
          "href": "https://api.example.org/v3/security_groups/guid-1"
        }
      }
    },
    {
      "guid": "guid-2",
      "name": "my-group2",
      "globally_enabled": {
        "running": false,
        "staging": true
      },
      "rules": [
        {
          "protocol": "tcp",
          "destination": "1.2.3.4/14",
          "ports": "443,80,8080"
        },
        {
          "protocol": "icmp",
          "destination": "1.2.3.4/16",
          "type": 5,
          "code": 0,
          "description": "test-desc-2"
        }
      ],
      "relationships": {
        "staging_spaces": {
          "data": [
            { "guid": "space-guid-3" },
            { "guid": "space-guid-4" }
          ]
        },
        "running_spaces": {
          "data": [
            { "guid": "space-guid-5" },
            { "guid": "space-guid-6" }
          ]
        }
      },
      "links": {
        "self": {
          "href": "https://api.example.org/v3/security_groups/guid-2"
        }
      }
    }
  ]
}`

const listV3SecurityGroupsByGuidPayload = `{
  "pagination": {
    "total_results": 1,
    "total_pages": 1,
    "first": {
      "href": "https://api.example.org/v3/security_groups?page=1&per_page=50"
    },
    "last": {
      "href": "https://api.example.org/v3/security_groups?page=1&per_page=50"
    },
    "next": null,
    "previous": null
  },
  "resources": [
    {
      "guid": "guid-1",
      "name": "my-group1",
      "globally_enabled": {
        "running": true,
        "staging": false
      },
      "rules": [
        {
          "protocol": "tcp",
          "destination": "1.2.3.4/10",
          "ports": "443,80,8080"
        },
        {
          "protocol": "icmp",
          "destination": "1.2.3.4/12",
          "type": 8,
          "code": 0,
          "description": "test-desc-1"
        }
      ],
      "relationships": {
        "staging_spaces": {
          "data": [
            { "guid": "space-guid-1" },
            { "guid": "space-guid-2" }
          ]
        },
        "running_spaces": {
          "data": []
        }
      },
      "links": {
        "self": {
          "href": "https://api.example.org/v3/security_groups/guid-1"
        }
      }
    }
  ]
}`

const genericV3SecurityGroupPayload = `{
  "guid": "guid-1",
  "name": "my-sec-group",
  "globally_enabled": {
    "running": true,
    "staging": false
  },
  "rules": [
    {
      "protocol": "tcp",
      "destination": "10.10.10.0/24",
      "ports": "443,80,8080"
    },
    {
      "protocol": "icmp",
      "destination": "10.10.11.0/24",
      "type": 8,
      "code": 0,
      "description": "Allow ping requests to private services"
    }
  ],
  "relationships": {
    "staging_spaces": {
      "data": []
    },
    "running_spaces": {
      "data": [
        { "guid": "space-guid-1" },
        { "guid": "space-guid-2" }
      ]
    }
  },
  "links": {
    "self": {
      "href": "https://api.example.org/v3/security_groups/guid-1"
    }
  }
}`

const setV3AppEnvironmentVariablesPayload = `{
  "var": {
    "RAILS_ENV": "production",
    "DEBUG": "false"
  },
  "links": {
    "self": {
      "href": "https://api.example.org/v3/apps/[guid]/environment_variables"
    },
    "app": {
      "href": "https://api.example.org/v3/apps/[guid]"
    }
  }
}`

const errorV3Payload = `{
  "errors": [
    {
      "code": 10008,
      "title": "CF-UnprocessableEntity",
      "detail": "something went wrong"
    }
  ]
}
`

const listV3DomainsPayload = `
  {
    "pagination": {
      "total_results": 1,
      "total_pages": 1,
      "first": {
        "href": "https://api.example.org/v3/domains?page=1&per_page=2"
      },
      "last": {
        "href": "https://api.example.org/v3/domains?page=2&per_page=2"
      },
      "next": null,
      "previous": null
    },
    "resources": [
      {
		"guid": "3a5d3d89-3f89-4f05-8188-8a2b298c79d5",
		"created_at": "2019-03-08T01:06:19Z",
		"updated_at": "2019-03-08T01:06:19Z",
		"name": "test-domain.com",
		"internal": false,
		"metadata": {
		  "labels": { },
		  "annotations": { }
		},
		"relationships": {
		  "organization": {
			"data": { "guid": "3a3f3d89-3f89-4f05-8188-751b298c79d5" }
		  },
		  "shared_organizations": {
			"data": [
			  {"guid": "404f3d89-3f89-6z72-8188-751b298d88d5"},
			  {"guid": "416d3d89-3f89-8h67-2189-123b298d3592"}
			]
		  }
		},
		"links": {
		  "self": {
			"href": "https://api.example.org/v3/domains/3a5d3d89-3f89-4f05-8188-8a2b298c79d5"
		  },
		  "organization": {
			"href": "https://api.example.org/v3/organizations/3a3f3d89-3f89-4f05-8188-751b298c79d5"
		  },
		  "route_reservations": {
			"href": "https://api.example.org/v3/domains/3a5d3d89-3f89-4f05-8188-8a2b298c79d5/route_reservations"
		  },
		  "shared_organizations": {
			"href": "https://api.example.org/v3/domains/3a5d3d89-3f89-4f05-8188-8a2b298c79d5/relationships/shared_organizations"
		  }
		}
	  }
    ]
}`

const listV3OrganizationsPayload = `{
  "pagination": {
    "total_results": 2,
    "total_pages": 2,
    "first": {
      "href": "https://api.example.org/v3/organizations?page=1&per_page=1"
    },
    "last": {
      "href": "https://api.example.org/v3/organizations?page=2&per_page=1"
    },
    "next": {
      "href": "https://api.example.org/v3/organizations?page=2&per_page=1"
    },
    "previous": null
  },
  "resources": [
    {
      "guid": "org-guid",
      "created_at": "2017-02-01T01:33:58Z",
      "updated_at": "2017-02-01T01:33:58Z",
      "name": "my-org-1",
      "relationships": {
        "quota": {
          "data": {
            "guid": "quota-guid"
          }
        }
      },
      "links": {
        "self": {
          "href": "https://api.example.org/v3/organizations/org-guid"
        },
        "domains": {
          "href": "https://api.example.org/v3/organizations/org-guid/domains"
        },
        "default_domain": {
          "href": "https://api.example.org/v3/organizations/org-guid/domains/default"
        },
        "quota": {
          "href": "https://api.example.org/v3/organization_quotas/quota-guid"
        }
      },
      "metadata": {
        "labels": {},
        "annotations": {}
      }
    }
  ]
}`

const listV3OrganizationsPayloadPage2 = `{
  "pagination": {
    "total_results": 2,
    "total_pages": 2,
    "first": {
      "href": "https://api.example.org/v3/organizations?page=1&per_page=1"
    },
    "last": {
      "href": "https://api.example.org/v3/organizations?page=2&per_page=1"
    },
    "next": null,
    "previous": {
      "href": "https://api.example.org/v3/organizations?page=2&per_page=1"
    }
  },
  "resources": [
    {
      "guid": "org-guid-2",
      "created_at": "2017-02-01T01:33:58Z",
      "updated_at": "2017-02-01T01:33:58Z",
      "name": "my-org-2",
      "relationships": {
        "quota": {
          "data": null
        }
      },
      "links": {
        "self": {
          "href": "https://api.example.org/v3/organizations/org-guid-2"
        },
        "domains": {
          "href": "https://api.example.org/v3/organizations/org-guid-2/domains"
        },
        "default_domain": {
          "href": "https://api.example.org/v3/organizations/org-guid-2/domains/default"
        }
      },
      "metadata": {
        "labels": {},
        "annotations": {}
      }
    }
  ]
}`

const updateV3OrganizationPayload = `
{
  "guid": "org-guid",
  "created_at": "2017-02-01T01:33:58Z",
  "updated_at": "2017-02-01T01:33:58Z",
  "name": "my-org",
  "suspended": false,
  "relationships": {
    "quota": {
      "data": null
    }
  },
  "links": {
    "self": {
      "href": "https://api.example.org/v3/organizations/org-guid"
    },
    "domains": {
      "href": "https://api.example.org/v3/organizations/org-guid/domains"
    },
    "default_domain": {
      "href": "https://api.example.org/v3/organizations/org-guid/domains/default"
    }
  },
  "metadata": {
    "labels": {
      "ORG_KEY": "org_value"
    },
    "annotations": {}
  }
}`

const getV3OrganizationPayload = `{
  "guid": "org-guid",
  "created_at": "2017-02-01T01:33:58Z",
  "updated_at": "2017-02-01T01:33:58Z",
  "name": "my-org",
  "relationships": {
    "quota": {
      "data": {
        "guid": "quota-guid"
      }
    }
  },
  "links": {
    "self": {
      "href": "https://api.example.org/v3/organizations/org-guid"
    },
    "domains": {
      "href": "https://api.example.org/v3/organizations/org-guid/domains"
    },
    "default_domain": {
      "href": "https://api.example.org/v3/organizations/org-guid/domains/default"
    }
  },
  "metadata": {
    "labels": {
      "ORG_KEY": "org_value"
    },
    "annotations": {}
  }
}`

const createV3OrganizationPayload = `{
  "guid": "org-guid",
  "created_at": "2017-02-01T01:33:58Z",
  "updated_at": "2017-02-01T01:33:58Z",
  "name": "my-org",
  "relationships": {
    "quota": {
      "data": {
        "guid": "quota-guid"
      }
    }
  },
  "links": {
    "self": {
      "href": "https://api.example.org/v3/organizations/org-guid"
    },
    "domains": {
      "href": "https://api.example.org/v3/organizations/org-guid/domains"
    },
    "default_domain": {
      "href": "https://api.example.org/v3/organizations/org-guid/domains/default"
    }
  },
  "metadata": {
    "labels": {
      "ORG_KEY": "org_value"
    },
    "annotations": {}
  }
}`

const listV3SpacesPayload = `{
  "pagination": {
    "total_results": 2,
    "total_pages": 2,
    "first": {
      "href": "https://api.example.org/v3/spaces?page=1&per_page=1"
    },
    "last": {
      "href": "https://api.example.org/v3/spaces?page=2&per_page=1"
    },
    "next": {
      "href": "https://api.example.org/v3/spaces?page=2&per_page=1"
    },
    "previous": null
  },
  "resources": [
    {
      "guid": "space-guid",
      "created_at": "2017-02-01T01:33:58Z",
      "updated_at": "2017-02-01T01:33:58Z",
      "name": "my-space-1",
      "relationships": {
        "organization": {
          "data": {
            "guid": "org-guid"
          }
        },
        "quota": {
          "data": null
        }
      },
      "links": {
        "self": {
          "href": "https://api.example.org/v3/spaces/space-guid"
        },
        "features": {
          "href": "https://api.example.org/v3/spaces/space-guid/features"
        },
        "organization": {
          "href": "https://api.example.org/v3/organizations/org-guid"
        },
        "apply_manifest": {
          "href": "https://api.example.org/v3/spaces/space-guid/actions/apply_manifest",
          "method": "POST"
        }
      },
      "metadata": {
        "labels": {},
        "annotations": {}
      }
    }
  ]
}`

const listV3SpacesPayloadPage2 = `{
  "pagination": {
    "total_results": 2,
    "total_pages": 2,
    "first": {
      "href": "https://api.example.org/v3/spaces?page=1&per_page=1"
    },
    "last": {
      "href": "https://api.example.org/v3/spaces?page=2&per_page=1"
    },
    "next": null,
    "previous": {
      "href": "https://api.example.org/v3/spaces?page=2&per_page=1"
    }
  },
  "resources": [
    {
      "guid": "space-guid-2",
      "created_at": "2017-02-01T01:33:58Z",
      "updated_at": "2017-02-01T01:33:58Z",
      "name": "my-space-2",
      "relationships": {
        "organization": {
          "data": {
            "guid": "org-guid"
          }
        },
        "quota": {
          "data": null
        }
      },
      "links": {
        "self": {
          "href": "https://api.example.org/v3/spaces/space-guid-2"
        },
        "features": {
          "href": "https://api.example.org/v3/spaces/space-guid-2/features"
        },
        "organization": {
          "href": "https://api.example.org/v3/organizations/org-guid"
        },
        "apply_manifest": {
          "href": "https://api.example.org/v3/spaces/space-guid-2/actions/apply_manifest",
          "method": "POST"
        }
      },
      "metadata": {
        "labels": {},
        "annotations": {}
      }
    }
  ]
}`

const listV3SpaceUsersPayload = `{
  "pagination": {
    "total_results": 2,
    "total_pages": 2,
    "first": {
      "href": "https://api.example.org/v3/spaces/space-guid/users?page=1&per_page=1"
    },
    "last": {
      "href": "https://api.example.org/v3/spaces/space-guid/users?page=2&per_page=1"
    },
    "next": {
      "href": "https://api.example.org/v3/spaces/space-guid/users?page=2&per_page=1"
    },
    "previous": null
  },
  "resources": [
    {
      "guid": "10a93b89-3f89-4f05-7238-8a2b123c79l9",
      "created_at": "2019-03-08T01:06:18Z",
      "updated_at": "2019-03-08T01:06:18Z",
      "username": "some-name-1",
      "presentation_name": "some-name-1",
      "origin": "uaa",
      "metadata": {
        "labels": {},
        "annotations":{}
      },
      "links": {
        "self": {
          "href": "https://api.example.org/v3/users/10a93b89-3f89-4f05-7238-8a2b123c79l9"
        }
      }
    }
  ]
}`

const listV3SpaceUsersPayloadPage2 = `{
  "pagination": {
    "total_results": 2,
    "total_pages": 2,
    "first": {
      "href": "https://api.example.org/v3/spaces/space-guid/users?page=1&per_page=1"
    },
    "last": {
      "href": "https://api.example.org/v3/spaces/space-guid/users?page=2&per_page=1"
    },
    "next": null,
    "previous": {
      "href": "https://api.example.org/v3/spaces/space-guid/users?page=1&per_page=1"
    }
  },
  "resources": [
    {
      "guid": "9da93b89-3f89-4f05-7238-8a2b123c79l9",
      "created_at": "2019-03-08T01:06:19Z",
      "updated_at": "2019-03-08T01:06:19Z",
      "username": "some-name-2",
      "presentation_name": "some-name-2",
      "origin": "ldap",
      "metadata": {
        "labels": {},
        "annotations":{}
      },
      "links": {
        "self": {
          "href": "https://api.example.org/v3/users/9da93b89-3f89-4f05-7238-8a2b123c79l9"
        }
      }
    }
  ]
}`

const updateV3SpacePayload = `
{
  "guid": "space-guid",
  "created_at": "2017-02-01T01:33:58Z",
  "updated_at": "2017-02-01T01:33:58Z",
  "name": "my-space",
  "relationships": {
    "organization": {
      "data": {
        "guid": "org-guid"
      }
    },
    "quota": {
      "data": null
    }
  },
  "links": {
    "self": {
      "href": "https://api.example.org/v3/spaces/space-guid"
    },
    "features": {
      "href": "https://api.example.org/v3/spaces/space-guid/features"
    },
    "organization": {
      "href": "https://api.example.org/v3/organizations/org-guid"
    },
    "apply_manifest": {
      "href": "https://api.example.org/v3/spaces/space-guid/actions/apply_manifest",
      "method": "POST"
    }
  },
  "metadata": {
    "labels": {
      "SPACE_KEY": "space_value"
    },
    "annotations": {}
  }
}`

const getV3SpacePayload = `{
  "guid": "space-guid",
  "created_at": "2017-02-01T01:33:58Z",
  "updated_at": "2017-02-01T01:33:58Z",
  "name": "my-space",
  "relationships": {
    "organization": {
      "data": {
        "guid": "org-guid"
      }
    },
    "quota": {
      "data": null
    }
  },
  "links": {
    "self": {
      "href": "https://api.example.org/v3/spaces/space-guid"
    },
    "features": {
      "href": "https://api.example.org/v3/spaces/space-guid/features"
    },
    "organization": {
      "href": "https://api.example.org/v3/organizations/org-guid"
    },
    "apply_manifest": {
      "href": "https://api.example.org/v3/spaces/space-guid/actions/apply_manifest",
      "method": "POST"
    }
  },
  "metadata": {
    "labels": {
      "SPACE_KEY": "space_value"
    },
    "annotations": {}
  }
}`

const createV3RoutePayload = `{
	"guid": "cbad697f-cac1-48f4-9017-ac08f39dfb31",
	"host": "a-hostname",
	"path": "/some_path",
	"url": "a-hostname.a-domain.com/some_path",
	"created_at": "2019-05-10T17:17:48Z",
	"updated_at": "2019-05-10T17:17:48Z",
	"metadata": {
	  "labels": { "key": "value" },
	  "annotations": { "note": "detailed information" }
	},
	"relationships": {
	  "space": {
		"data": {
		  "guid": "885a8cb3-c07b-4856-b448-eeb10bf36236"
		}
	  },
	  "domain": {
		"data": {
		  "guid": "0b5f3633-194c-42d2-9408-972366617e0e"
		}
	  }
	},
	"links": {
	  "self": {
		"href": "https://api.example.org/v3/routes/cbad697f-cac1-48f4-9017-ac08f39dfb31"
	  },
	  "space": {
		"href": "https://api.example.org/v3/spaces/885a8cb3-c07b-4856-b448-eeb10bf36236"
	  },
	  "domain": {
		"href": "https://api.example.org/v3/domains/0b5f3633-194c-42d2-9408-972366617e0e"
	  },
	  "destinations": {
		"href": "https://api.example.org/v3/routes/cbad697f-cac1-48f4-9017-ac08f39dfb31/destinations"
	  }
	}
}`

const createV3SpaceRolePayload = `{
   "guid": "b9f59ab2-2b09-438e-bebb-30e8704ffb89",
   "created_at": "2022-05-31T20:14:13Z",
   "updated_at": "2022-05-31T20:14:13Z",
   "type": "space_supporter",
   "relationships": {
      "user": {
         "data": {
            "guid": "c4958204-6b65-43ea-832b-e4c57aea6641"
         }
      },
      "space": {
         "data": {
            "guid": "b40a40c8-58b7-49a0-b47d-9d6fe5d72905"
         }
      },
      "organization": {
         "data": null
      }
   },
   "links": {
      "self": {
         "href": "https://api.example.org/v3/roles/b9f59ab2-2b09-438e-bebb-30e8704ffb89"
      },
      "user": {
         "href": "https://api.example.org/v3/users/c4958204-6b65-43ea-832b-e4c57aea6641"
      },
      "space": {
         "href": "https://api.example.org/v3/spaces/b40a40c8-58b7-49a0-b47d-9d6fe5d72905"
      }
   }
}`

const createV3OrganizationRolePayload = `{
  "guid": "21cbfaeb-bff7-4cfd-a7a9-6c13ec76f246",
  "created_at": "2022-05-31T18:19:42Z",
  "updated_at": "2022-05-31T18:19:42Z",
  "type": "organization_user",
  "relationships": {
    "user": {
      "data": {
        "guid": "ac2e02c9-2c5c-4712-a620-a68449d263c3"
      }
    },
    "organization": {
      "data": {
        "guid": "fa8a8346-0d92-4729-870c-77ee1934f973"
      }
    },
    "space": {
      "data": null
    }
  },
  "links": {
    "self": {
      "href": "https://api.example.org/v3/roles/21cbfaeb-bff7-4cfd-a7a9-6c13ec76f246"
    },
    "user": {
      "href": "https://api.example.org/v3/users/ac2e02c9-2c5c-4712-a620-a68449d263c3"
    },
    "organization": {
      "href": "https://api.example.org/v3/organizations/fa8a8346-0d92-4729-870c-77ee1934f973"
    }
  }
}`

const createV3SpacePayload = `{
  "guid": "space-guid",
  "created_at": "2017-02-01T01:33:58Z",
  "updated_at": "2017-02-01T01:33:58Z",
  "name": "my-space",
  "relationships": {
    "organization": {
      "data": {
        "guid": "org-guid"
      }
    },
    "quota": {
      "data": null
    }
  },
  "links": {
    "self": {
      "href": "https://api.example.org/v3/spaces/space-guid"
    },
    "features": {
      "href": "https://api.example.org/v3/spaces/space-guid/features"
    },
    "organization": {
      "href": "https://api.example.org/v3/organizations/org-guid"
    },
    "apply_manifest": {
      "href": "https://api.example.org/v3/spaces/space-guid/actions/apply_manifest",
      "method": "POST"
    }
  },
  "metadata": {
    "labels": {
      "SPACE_KEY": "space_value"
    },
    "annotations": {}
  }
}`

const getV3AppPayload = `{
  "guid": "1cb006ee-fb05-47e1-b541-c34179ddc446",
  "name": "my_app",
  "state": "STOPPED",
  "created_at": "2016-03-17T21:41:30Z",
  "updated_at": "2016-06-08T16:41:26Z",
  "lifecycle": {
    "type": "buildpack",
    "data": {
      "buildpacks": ["java_buildpack"],
      "stack": "cflinuxfs2"
    }
  },
  "relationships": {
    "space": {
      "data": {
        "guid": "2f35885d-0c9d-4423-83ad-fd05066f8576"
      }
    }
  },
  "links": {
    "self": {
      "href": "https://api.example.org/v3/apps/1cb006ee-fb05-47e1-b541-c34179ddc446"
    },
    "space": {
      "href": "https://api.example.org/v3/spaces/2f35885d-0c9d-4423-83ad-fd05066f8576"
    },
    "processes": {
      "href": "https://api.example.org/v3/apps/1cb006ee-fb05-47e1-b541-c34179ddc446/processes"
    },
    "route_mappings": {
      "href": "https://api.example.org/v3/apps/1cb006ee-fb05-47e1-b541-c34179ddc446/route_mappings"
    },
    "packages": {
      "href": "https://api.example.org/v3/apps/1cb006ee-fb05-47e1-b541-c34179ddc446/packages"
    },
    "environment_variables": {
      "href": "https://api.example.org/v3/apps/1cb006ee-fb05-47e1-b541-c34179ddc446/environment_variables"
    },
    "current_droplet": {
      "href": "https://api.example.org/v3/apps/1cb006ee-fb05-47e1-b541-c34179ddc446/droplets/current"
    },
    "droplets": {
      "href": "https://api.example.org/v3/apps/1cb006ee-fb05-47e1-b541-c34179ddc446/droplets"
    },
    "tasks": {
      "href": "https://api.example.org/v3/apps/1cb006ee-fb05-47e1-b541-c34179ddc446/tasks"
    },
    "start": {
      "href": "https://api.example.org/v3/apps/1cb006ee-fb05-47e1-b541-c34179ddc446/actions/start",
      "method": "POST"
    },
    "stop": {
      "href": "https://api.example.org/v3/apps/1cb006ee-fb05-47e1-b541-c34179ddc446/actions/stop",
      "method": "POST"
    }
  },
  "metadata": {
    "labels": {},
    "annotations": {
      "contacts": "Bill tel(1111111) email(bill@fixme), Bob tel(222222) pager(3333333#555) email(bob@fixme)"
    }
  }
}`

const getV3AppEnvPayload = `{
  "staging_env_json": {
    "GEM_CACHE": "http://gem-cache.example.org"
  },
  "running_env_json": {
    "HTTP_PROXY": "http://proxy.example.org"
  },
  "environment_variables": {
    "RAILS_ENV": "production"
  },
  "system_env_json": {
    "VCAP_SERVICES": {
      "mysql": [
        {
          "name": "db-for-my-app",
          "label": "mysql",
          "tags": ["relational", "sql"],
          "plan": "xlarge",
          "credentials": {
            "username": "user",
            "password": "top-secret"
           },
          "syslog_drain_url": "https://syslog.example.org/drain",
          "provider": null
        }
      ]
    }
  },
  "application_env_json": {
    "VCAP_APPLICATION": {
      "limits": {
        "fds": 16384
      },
      "application_name": "my_app",
      "application_uris": [ "my_app.example.org" ],
      "name": "my_app",
      "space_name": "my_space",
      "space_id": "2f35885d-0c9d-4423-83ad-fd05066f8576",
      "uris": [ "my_app.example.org" ],
      "users": null
    }
  }
}`

const createV3AppPayload = `{
  "guid": "app-guid",
  "name": "my-app",
  "state": "STOPPED",
  "created_at": "2016-03-17T21:41:30Z",
  "updated_at": "2016-06-08T16:41:26Z",
  "lifecycle": {
    "type": "buildpack",
    "data": {
      "buildpacks": ["java_buildpack"],
      "stack": "cflinuxfs2"
    }
  },
  "relationships": {
    "space": {
      "data": {
        "guid": "space-guid"
      }
    }
  },
  "links": {
    "self": {
      "href": "https://api.example.org/v3/apps/app-guid"
    },
    "space": {
      "href": "https://api.example.org/v3/spaces/space-guid"
    },
    "processes": {
      "href": "https://api.example.org/v3/apps/app-guid/processes"
    },
    "route_mappings": {
      "href": "https://api.example.org/v3/apps/app-guid/route_mappings"
    },
    "packages": {
      "href": "https://api.example.org/v3/apps/app-guid/packages"
    },
    "environment_variables": {
      "href": "https://api.example.org/v3/apps/app-guid/environment_variables"
    },
    "current_droplet": {
      "href": "https://api.example.org/v3/apps/app-guid/droplets/current"
    },
    "droplets": {
      "href": "https://api.example.org/v3/apps/app-guid/droplets"
    },
    "tasks": {
      "href": "https://api.example.org/v3/apps/app-guid/tasks"
    },
    "start": {
      "href": "https://api.example.org/v3/apps/app-guid/actions/start",
      "method": "POST"
    },
    "stop": {
      "href": "https://api.example.org/v3/apps/app-guid/actions/stop",
      "method": "POST"
    }
  },
  "metadata": {
    "labels": {},
    "annotations": {}
  }
}`

const createV3BuildPayload = `{
  "guid": "585bc3c1-3743-497d-88b0-403ad6b56d16",
  "created_at": "2016-03-28T23:39:34Z",
  "updated_at": "2016-06-08T16:41:26Z",
  "created_by": {
    "guid": "3cb4e243-bed4-49d5-8739-f8b45abdec1c",
    "name": "bill",
    "email": "bill@example.com"
  },
  "state": "STAGING",
  "error": null,
  "lifecycle": {
    "type": "buildpack",
    "data": {
      "buildpacks": [ "ruby_buildpack" ],
      "stack": "cflinuxfs2"
    }
  },
  "package": {
    "guid": "8e4da443-f255-499c-8b47-b3729b5b7432"
  },
  "droplet": null,
  "metadata": {
    "labels": { },
    "annotations": { }
  },
  "links": {
    "self": {
      "href": "https://api.example.org/v3/builds/585bc3c1-3743-497d-88b0-403ad6b56d16"
    },
    "app": {
      "href": "https://api.example.org/v3/apps/7b34f1cf-7e73-428a-bb5a-8a17a8058396"
    }
  }
}
`

const startV3AppPayload = `{
  "guid": "1cb006ee-fb05-47e1-b541-c34179ddc446",
  "name": "my_app",
  "state": "STARTED",
  "created_at": "2016-03-17T21:41:30Z",
  "updated_at": "2016-03-18T11:32:30Z",
  "lifecycle": {
    "type": "buildpack",
    "data": {
      "buildpacks": ["java_buildpack"],
      "stack": "cflinuxfs2"
    }
  },
  "relationships": {
    "space": {
      "data": {
        "guid": "2f35885d-0c9d-4423-83ad-fd05066f8576"
      }
    }
  },
  "links": {
    "self": {
      "href": "https://api.example.org/v3/apps/1cb006ee-fb05-47e1-b541-c34179ddc446"
    },
    "space": {
      "href": "https://api.example.org/v3/spaces/2f35885d-0c9d-4423-83ad-fd05066f8576"
    },
    "processes": {
      "href": "https://api.example.org/v3/apps/1cb006ee-fb05-47e1-b541-c34179ddc446/processes"
    },
    "route_mappings": {
      "href": "https://api.example.org/v3/apps/1cb006ee-fb05-47e1-b541-c34179ddc446/route_mappings"
    },
    "packages": {
      "href": "https://api.example.org/v3/apps/1cb006ee-fb05-47e1-b541-c34179ddc446/packages"
    },
    "environment_variables": {
      "href": "https://api.example.org/v3/apps/1cb006ee-fb05-47e1-b541-c34179ddc446/environment_variables"
    },
    "current_droplet": {
      "href": "https://api.example.org/v3/apps/1cb006ee-fb05-47e1-b541-c34179ddc446/droplets/current"
    },
    "droplets": {
      "href": "https://api.example.org/v3/apps/1cb006ee-fb05-47e1-b541-c34179ddc446/droplets"
    },
    "tasks": {
      "href": "https://api.example.org/v3/apps/1cb006ee-fb05-47e1-b541-c34179ddc446/tasks"
    },
    "start": {
      "href": "https://api.example.org/v3/apps/1cb006ee-fb05-47e1-b541-c34179ddc446/actions/start",
      "method": "POST"
    },
    "stop": {
      "href": "https://api.example.org/v3/apps/1cb006ee-fb05-47e1-b541-c34179ddc446/actions/stop",
      "method": "POST"
    }
  },
  "metadata": {
    "labels": {},
    "annotations": {}
  }
}`

const updateV3AppPayload = `{
  "guid": "1cb006ee-fb05-47e1-b541-c34179ddc446",
  "name": "my_app",
  "state": "STARTED",
  "created_at": "2016-03-17T21:41:30Z",
  "updated_at": "2016-03-18T11:32:30Z",
  "lifecycle": {
    "type": "buildpack",
    "data": {
      "buildpacks": ["java_buildpack"],
      "stack": "cflinuxfs2"
    }
  },
  "relationships": {
    "space": {
      "data": {
        "guid": "2f35885d-0c9d-4423-83ad-fd05066f8576"
      }
    }
  },
  "links": {
    "self": {
      "href": "https://api.example.org/v3/apps/1cb006ee-fb05-47e1-b541-c34179ddc446"
    },
    "space": {
      "href": "https://api.example.org/v3/spaces/2f35885d-0c9d-4423-83ad-fd05066f8576"
    },
    "processes": {
      "href": "https://api.example.org/v3/apps/1cb006ee-fb05-47e1-b541-c34179ddc446/processes"
    },
    "route_mappings": {
      "href": "https://api.example.org/v3/apps/1cb006ee-fb05-47e1-b541-c34179ddc446/route_mappings"
    },
    "packages": {
      "href": "https://api.example.org/v3/apps/1cb006ee-fb05-47e1-b541-c34179ddc446/packages"
    },
    "environment_variables": {
      "href": "https://api.example.org/v3/apps/1cb006ee-fb05-47e1-b541-c34179ddc446/environment_variables"
    },
    "current_droplet": {
      "href": "https://api.example.org/v3/apps/1cb006ee-fb05-47e1-b541-c34179ddc446/droplets/current"
    },
    "droplets": {
      "href": "https://api.example.org/v3/apps/1cb006ee-fb05-47e1-b541-c34179ddc446/droplets"
    },
    "tasks": {
      "href": "https://api.example.org/v3/apps/1cb006ee-fb05-47e1-b541-c34179ddc446/tasks"
    },
    "start": {
      "href": "https://api.example.org/v3/apps/1cb006ee-fb05-47e1-b541-c34179ddc446/actions/start",
      "method": "POST"
    },
    "stop": {
      "href": "https://api.example.org/v3/apps/1cb006ee-fb05-47e1-b541-c34179ddc446/actions/stop",
      "method": "POST"
    }
  },
  "metadata": {
    "labels": {
      "environment": "production",
      "internet-facing": "false"
    },
    "annotations": {}
  }
}`

const currentDropletV3AppPayload = `{
  "data": {
    "guid": "9d8e007c-ce52-4ea7-8a57-f2825d2c6b39"
  },
  "links": {
    "self": {
      "href": "https://api.example.org/v3/apps/d4c91047-7b29-4fda-b7f9-04033e5c9c9f/relationships/current_droplet"
    },
    "related": {
      "href": "https://api.example.org/v3/apps/d4c91047-7b29-4fda-b7f9-04033e5c9c9f/droplets/current"
    }
  }
}`

const getV3CurrentAppDropletPayload = `{
  "guid": "585bc3c1-3743-497d-88b0-403ad6b56d16",
  "state": "STAGED",
  "error": null,
  "lifecycle": {
    "type": "buildpack",
    "data": {}
  },
  "execution_metadata": "",
  "process_types": {
    "rake": "bundle exec rake",
    "web": "bundle exec rackup config.ru -p $PORT"
  },
  "checksum": {
    "type": "sha256",
    "value": "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
  },
  "buildpacks": [
    {
      "name": "ruby_buildpack",
      "detect_output": "ruby 1.6.14",
      "version": "1.1.1.",
      "buildpack_name": "ruby"
    }
  ],
  "stack": "cflinuxfs3",
  "image": null,
  "created_at": "2016-03-28T23:39:34Z",
  "updated_at": "2016-03-28T23:39:47Z",
  "relationships": {
    "app": {
      "data": {
        "guid": "7b34f1cf-7e73-428a-bb5a-8a17a8058396"
      }
    }
  },
  "links": {
    "self": {
      "href": "https://api.example.org/v3/droplets/585bc3c1-3743-497d-88b0-403ad6b56d16"
    },
    "package": {
      "href": "https://api.example.org/v3/packages/8222f76a-9e09-4360-b3aa-1ed329945e92"
    },
    "app": {
      "href": "https://api.example.org/v3/apps/7b34f1cf-7e73-428a-bb5a-8a17a8058396"
    },
    "assign_current_droplet": {
      "href": "https://api.example.org/v3/apps/7b34f1cf-7e73-428a-bb5a-8a17a8058396/relationships/current_droplet",
      "method": "PATCH"
      },
    "download": {
      "href": "https://api.example.org/v3/droplets/585bc3c1-3743-497d-88b0-403ad6b56d16/download"
    }
  },
  "metadata": {
    "labels": {},
    "annotations": {}
  }
}`

const listV3AppsPayload = `{
  "pagination": {
    "total_results": 2,
    "total_pages": 2,
    "first": {
      "href": "https://api.example.org/v3/apps?page=1&per_page=2"
    },
    "last": {
      "href": "https://api.example.org/v3/apps?page=2&per_page=2"
    },
    "next": {
      "href": "https://api.example.org/v3/apps?page=2&per_page=2"
    },
    "previous": null
  },
  "resources": [
    {
      "guid": "1cb006ee-fb05-47e1-b541-c34179ddc446",
      "name": "my_app",
      "state": "STARTED",
      "created_at": "2016-03-17T21:41:30Z",
      "updated_at": "2016-03-18T11:32:30Z",
      "lifecycle": {
        "type": "buildpack",
        "data": {
          "buildpacks": ["java_buildpack"],
          "stack": "cflinuxfs2"
        }
      },
      "relationships": {
        "space": {
          "data": {
            "guid": "2f35885d-0c9d-4423-83ad-fd05066f8576"
          }
        }
      },
      "links": {
        "self": {
          "href": "https://api.example.org/v3/apps/1cb006ee-fb05-47e1-b541-c34179ddc446"
        },
        "space": {
          "href": "https://api.example.org/v3/spaces/2f35885d-0c9d-4423-83ad-fd05066f8576"
        },
        "processes": {
          "href": "https://api.example.org/v3/apps/1cb006ee-fb05-47e1-b541-c34179ddc446/processes"
        },
        "route_mappings": {
          "href": "https://api.example.org/v3/apps/1cb006ee-fb05-47e1-b541-c34179ddc446/route_mappings"
        },
        "packages": {
          "href": "https://api.example.org/v3/apps/1cb006ee-fb05-47e1-b541-c34179ddc446/packages"
        },
        "environment_variables": {
          "href": "https://api.example.org/v3/apps/1cb006ee-fb05-47e1-b541-c34179ddc446/environment_variables"
        },
        "current_droplet": {
          "href": "https://api.example.org/v3/apps/1cb006ee-fb05-47e1-b541-c34179ddc446/droplets/current"
        },
        "droplets": {
          "href": "https://api.example.org/v3/apps/1cb006ee-fb05-47e1-b541-c34179ddc446/droplets"
        },
        "tasks": {
          "href": "https://api.example.org/v3/apps/1cb006ee-fb05-47e1-b541-c34179ddc446/tasks"
        },
        "start": {
          "href": "https://api.example.org/v3/apps/1cb006ee-fb05-47e1-b541-c34179ddc446/actions/start",
          "method": "POST"
        },
        "stop": {
          "href": "https://api.example.org/v3/apps/1cb006ee-fb05-47e1-b541-c34179ddc446/actions/stop",
          "method": "POST"
        }
      },
      "metadata": {
        "labels": {},
        "annotations": {}
      }
    }
  ]
}`

const listV3AppsPayloadPage2 = `{
  "pagination": {
    "total_results": 2,
    "total_pages": 2,
    "first": {
      "href": "https://api.example.org/v3/apps?page=1&per_page=2"
    },
    "last": {
      "href": "https://api.example.org/v3/apps?page=2&per_page=2"
    },
    "next": null,
    "previous": null
  },
  "resources": [
    {
      "guid": "02b4ec9b-94c7-4468-9c23-4e906191a0f8",
      "name": "my_app2",
      "state": "STOPPED",
      "created_at": "1970-01-01T00:00:02Z",
      "updated_at": "2016-06-08T16:41:26Z",
      "lifecycle": {
        "type": "buildpack",
        "data": {
          "buildpacks": ["ruby_buildpack", "staticfile_buildpack"],
          "stack": "cflinuxfs2"
        }
      },
      "relationships": {
        "space": {
          "data": {
            "guid": "2f35885d-0c9d-4423-83ad-fd05066f8576"
          }
        }
      },
      "links": {
        "self": {
          "href": "https://api.example.org/v3/apps/02b4ec9b-94c7-4468-9c23-4e906191a0f8"
        },
        "space": {
          "href": "https://api.example.org/v3/spaces/2f35885d-0c9d-4423-83ad-fd05066f8576"
        },
        "processes": {
          "href": "https://api.example.org/v3/apps/02b4ec9b-94c7-4468-9c23-4e906191a0f8/processes"
        },
        "route_mappings": {
          "href": "https://api.example.org/v3/apps/02b4ec9b-94c7-4468-9c23-4e906191a0f8/route_mappings"
        },
        "packages": {
          "href": "https://api.example.org/v3/apps/02b4ec9b-94c7-4468-9c23-4e906191a0f8/packages"
        },
        "environment_variables": {
          "href": "https://api.example.org/v3/apps/02b4ec9b-94c7-4468-9c23-4e906191a0f8/environment_variables"
        },
        "current_droplet": {
          "href": "https://api.example.org/v3/apps/02b4ec9b-94c7-4468-9c23-4e906191a0f8/droplets/current"
        },
        "droplets": {
          "href": "https://api.example.org/v3/apps/02b4ec9b-94c7-4468-9c23-4e906191a0f8/droplets"
        },
        "tasks": {
          "href": "https://api.example.org/v3/apps/02b4ec9b-94c7-4468-9c23-4e906191a0f8/tasks"
        },
        "start": {
          "href": "https://api.example.org/v3/apps/02b4ec9b-94c7-4468-9c23-4e906191a0f8/actions/start",
          "method": "POST"
        },
        "stop": {
          "href": "https://api.example.org/v3/apps/02b4ec9b-94c7-4468-9c23-4e906191a0f8/actions/stop",
          "method": "POST"
        }
      },
      "metadata": {
        "labels": {},
        "annotations": {}
      }
    }
  ]
}`

const listV3ServiceInstancesPayload = `{
  "pagination": {
    "total_results": 1,
    "total_pages": 1,
    "first": {
      "href": "https://api.example.org/v3/service_instances?page=1&per_page=50"
    },
    "last": {
      "href": "https://api.example.org/v3/service_instances?page=1&per_page=50"
    },
    "next": null,
    "previous": null
  },
  "resources": [
    {
      "guid": "85ccdcad-d725-4109-bca4-fd6ba062b5c8",
      "created_at": "2017-11-17T13:54:21Z",
      "updated_at": "2017-11-17T13:54:21Z",
      "name": "my_service_instance",
      "relationships": {
        "space": {
          "data": {
            "guid": "ae0031f9-dd49-461c-a945-df40e77c39cb"
          }
        }
      },
      "metadata": {
        "labels": { },
        "annotations": { }
      },
      "links": {
        "space": {
          "href": "https://api.example.org/v3/spaces/ae0031f9-dd49-461c-a945-df40e77c39cb"
        }
      }
    }
  ]
}`

const listV3RoutesPayload = `{
	"pagination": {
	  "total_results": 1,
	  "total_pages": 1,
	  "first": {
		"href": "https://api.example.org/v3/routes?page=1&per_page=1"
	  },
	  "last": {
		"href": "https://api.example.org/v3/routes?page=1&per_page=1"
	  },
	  "next": null,
	  "previous": null
	},
	"resources": [
	  {
		"guid": "cbad697f-cac1-48f4-9017-ac08f39dfb31",
		"host": "a-hostname",
		"path": "/some_path",
		"url": "a-hostname.a-domain.com/some_path",
		"created_at": "2019-05-10T17:17:48Z",
		"updated_at": "2019-05-10T17:17:48Z",
		"metadata": {
		  "labels": {},
		  "annotations": {}
		},
		"relationships": {
		  "space": {
			"data": {
			  "guid": "885a8cb3-c07b-4856-b448-eeb10bf36236"
			}
		  },
		  "domain": {
			"data": {
			  "guid": "0b5f3633-194c-42d2-9408-972366617e0e"
			}
		  }
		},
		"links": {
		  "self": {
			"href": "https://api.example.org/v3/routes/cbad697f-cac1-48f4-9017-ac08f39dfb31"
		  },
		  "space": {
			"href": "https://api.example.org/v3/spaces/885a8cb3-c07b-4856-b448-eeb10bf36236"
		  },
		  "domain": {
			"href": "https://api.example.org/v3/domains/0b5f3633-194c-42d2-9408-972366617e0e"
		  },
		  "destinations": {
			"href": "https://api.example.org/v3/routes/cbad697f-cac1-48f4-9017-ac08f39dfb31/destinations"
		  }
		}
	  }
	]
}`

const listPackagesForV3AppPayloadPage1 = `{
  "pagination": {
    "total_results": 2,
    "total_pages": 2,
    "first": {
      "href": "https://api.example.org/v3/apps/f2efe391-2b5b-4836-8518-ad93fa9ebf69/packages?page=1&per_page=1"
    },
    "last": {
      "href": "https://api.example.org/v3/apps/f2efe391-2b5b-4836-8518-ad93fa9ebf69/packages?page=2&per_page=1"
    },
    "next": {
      "href": "https://api.example.org/v3/apps/f2efe391-2b5b-4836-8518-ad93fa9ebf69/packages?page=2&per_page=1"
    },
    "previous": null
  },
  "resources": [
    {
      "guid": "752edab0-2147-4f58-9c25-cd72ad8c3561",
      "type": "bits",
      "data": {
        "error": null,
        "checksum": {
          "type": "sha256",
          "value": null
        }
      },
      "state": "READY",
      "created_at": "2016-03-17T21:41:09Z",
      "updated_at": "2016-06-08T16:41:26Z",
      "links": {
        "self": {
          "href": "https://api.example.org/v3/packages/752edab0-2147-4f58-9c25-cd72ad8c3561"
        },
        "upload": {
          "href": "https://api.example.org/v3/packages/752edab0-2147-4f58-9c25-cd72ad8c3561/upload",
          "method": "POST"
        },
        "download": {
          "href": "https://api.example.org/v3/packages/752edab0-2147-4f58-9c25-cd72ad8c3561/download",
          "method": "GET"
        },
        "app": {
          "href": "https://api.example.org/v3/apps/f2efe391-2b5b-4836-8518-ad93fa9ebf69"
        }
      },
      "metadata": {
        "labels": {},
        "annotations": {}
      }
    }
  ]
}`

const listPackagesForV3AppPayloadPage2 = `{
  "pagination": {
    "total_results": 2,
    "total_pages": 2,
    "first": {
      "href": "https://api.example.org/v3/apps/f2efe391-2b5b-4836-8518-ad93fa9ebf69/packages?page=1&per_page=1"
    },
    "last": {
      "href": "https://api.example.org/v3/apps/f2efe391-2b5b-4836-8518-ad93fa9ebf69/packages?page=2&per_page=1"
    },
    "next": null,
    "previous": null
  },
  "resources": [
    {
      "guid": "2345ab-2147-4f58-9c25-cd72ad8c3561",
      "type": "bits",
      "data": {
        "error": null,
        "checksum": {
          "type": "sha256",
          "value": null
        }
      },
      "state": "READY",
      "created_at": "2016-03-17T21:41:09Z",
      "updated_at": "2016-06-08T16:41:26Z",
      "links": {
        "self": {
          "href": "https://api.example.org/v3/packages/2345ab-2147-4f58-9c25-cd72ad8c3561"
        },
        "upload": {
          "href": "https://api.example.org/v3/packages/2345ab-2147-4f58-9c25-cd72ad8c3561/upload",
          "method": "POST"
        },
        "download": {
          "href": "https://api.example.org/v3/packages/2345ab-2147-4f58-9c25-cd72ad8c3561/download",
          "method": "GET"
        },
        "app": {
          "href": "https://api.example.org/v3/apps/f2efe391-2b5b-4836-8518-ad93fa9ebf69"
        }
      },
      "metadata": {
        "labels": {},
        "annotations": {}
      }
    }
  ]
}`

const copyPackageV3Payload = `{
  "guid": "fec72fc1-e453-4463-a86d-5df426f337a3",
  "type": "docker",
  "data": {
    "image": "http://awesome-sauce.example.org"
  },
  "state": "COPYING",
  "created_at": "2016-03-17T21:41:09Z",
  "updated_at": "2016-06-08T16:41:26Z",
  "links": {
    "self": {
      "href": "https://api.example.org/v3/packages/fec72fc1-e453-4463-a86d-5df426f337a3"
    },
    "app": {
      "href": "https://api.example.org/v3/apps/36208a68-562d-4f51-94ea-28bd8553a271"
    }
  },
  "metadata": {
    "labels": {},
    "annotations": {}
  }
}`

const getV3DeploymentPayload = `{
  "guid": "59c3d133-2b83-46f3-960e-7765a129aea4",
  "state": "DEPLOYING",
  "status": {
    "value": "ACTIVE",
    "reason": "DEPLOYING",
    "details": {
      "last_successful_healthcheck": "2018-04-25T22:42:10Z"
    }
  },
  "strategy": "rolling",
  "droplet": {
    "guid": "44ccfa61-dbcf-4a0d-82fe-f668e9d2a962"
  },
  "previous_droplet": {
    "guid": "cc6bc315-bd06-49ce-92c2-bc3ad45268c2"
  },
  "new_processes": [
    {
      "guid": "fd5d3e60-f88c-4c37-b1ae-667cfc65a856",
      "type": "web"
    }
  ],
  "revision": {
    "guid": "56126cba-656a-4eba-a81e-7e9951b2df57",
    "version": 1
  },
  "created_at": "2018-04-25T22:42:10Z",
  "updated_at": "2018-04-25T22:42:10Z",
  "metadata": {
    "labels": { },
    "annotations": { }
  },
  "relationships": {
    "app": {
      "data": {
        "guid": "305cea31-5a44-45ca-b51b-e89c7a8ef8b2"
      }
    }
  },
  "links": {
    "self": {
      "href": "https://api.example.org/v3/deployments/59c3d133-2b83-46f3-960e-7765a129aea4"
    },
    "app": {
      "href": "https://api.example.org/v3/apps/305cea31-5a44-45ca-b51b-e89c7a8ef8b2"
    }
  }
}`

const listV3StacksPayload = `{
  "pagination": {
    "total_results": 2,
    "total_pages": 1,
    "first": {
      "href": "https://api.example.org/v3/stacks?page=1&per_page=2"
    },
    "last": {
      "href": "https://api.example.org/v3/stacks?page=1&per_page=2"
    },
    "next": null,
    "previous": null
  },
  "resources": [
    {
      "guid": "guid-1",
      "created_at": "2018-11-09T22:43:28Z",
      "updated_at": "2018-11-09T22:43:28Z",
      "name": "my-stack-1",
      "description": "This is my first stack!",
      "metadata": {
        "labels": {},
        "annotations": {}
      },
      "links": {
        "self": {
          "href": "https://api.example.org/v3/stacks/guid-1"
        }
      }
    },
    {
      "guid": "guid-2",
      "created_at": "2018-11-09T22:43:29Z",
      "updated_at": "2018-11-09T22:43:29Z",
      "name": "my-stack-2",
      "description": "This is my second stack!",
      "metadata": {
        "labels": {},
        "annotations": {}
      },
      "links": {
        "self": {
          "href": "https://api.example.org/v3/stacks/guid-2"
        }
      }
    }
  ]
}`

const listV3ServiceCredentialBindingsPayload = `{
  "pagination": {
    "total_results": 1,
    "total_pages": 1,
    "first": {
      "href": "https://api.example.org/v3/service_instances?page=1&per_page=50"
    },
    "last": {
      "href": "https://api.example.org/v3/service_instances?page=1&per_page=50"
    },
    "next": null,
    "previous": null
  },
  "resources": [
    {
      "guid": "d9634934-8e1f-4c2d-bb33-fa5df019cf9d",
      "created_at": "2022-02-17T17:17:44Z",
      "updated_at": "2022-02-17T17:17:44Z",
      "name": "my_service_key",
      "type": "key",
      "relationships": {
        "service_instance": {
          "data": {
              "guid": "85ccdcad-d725-4109-bca4-fd6ba062b5c8"
          }
        }
      },
      "metadata": {
        "labels": { },
        "annotations": { }
      },
      "links": {
        "self": {
          "href": "https://api.example.org/v3/service_credential_bindings/d9634934-8e1f-4c2d-bb33-fa5df019cf9d"
        },
        "details": {
          "href": "https://api.example.org/v3/service_credential_bindings/d9634934-8e1f-4c2d-bb33-fa5df019cf9d/details"
        },
        "service_instance": {
          "href": "https://api.example.org/v3/service_instances/85ccdcad-d725-4109-bca4-fd6ba062b5c8"
        },
        "parameters": {
          "href": "https://api.example.org/v3/service_credential_bindings/d9634934-8e1f-4c2d-bb33-fa5df019cf9d/parameters"
        }
      }
    }
  ]
}`

const GetV3ServiceCredentialBindingsByGUIDPayload = `{
  "guid": "d9634934-8e1f-4c2d-bb33-fa5df019cf9d",
  "created_at": "2022-02-17T17:17:44Z",
  "updated_at": "2022-02-17T17:17:44Z",
  "name": "my_service_key",
  "type": "key",
  "last_operation": {
    "type": "create",
    "state": "succeeded",
    "description": "",
    "created_at": "2022-02-17T17:17:44Z",
    "updated_at": "2022-02-17T17:17:44Z"
  },
  "relationships": {
    "service_instance": {
      "data": {
        "guid": "85ccdcad-d725-4109-bca4-fd6ba062b5c8"
      }
    }
  },
  "metadata": {
    "labels": { },
    "annotations": { }
  },
  "links": {
    "self": {
      "href": "https://api.example.org/v3/service_credential_bindings/d9634934-8e1f-4c2d-bb33-fa5df019cf9d"
    },
    "details": {
      "href": "https://api.example.org/v3/service_credential_bindings/d9634934-8e1f-4c2d-bb33-fa5df019cf9d/details"
    },
    "service_instance": {
      "href": "https://api.example.org/v3/service_instances/85ccdcad-d725-4109-bca4-fd6ba062b5c8"
    },
    "parameters": {
      "href": "https://api.example.org/v3/service_credential_bindings/d9634934-8e1f-4c2d-bb33-fa5df019cf9d/parameters"
    }
  }
}`

const listV3UsersPayload = `{
   "pagination": {
      "total_results": 3,
      "total_pages": 2,
      "first": {
         "href": "https://api.example.org/v3/users?page=1&per_page=2"
      },
      "last": {
         "href": "https://api.example.org/v3/users?page=2&per_page=2"
      },
      "next": {
         "href": "https://api.example.org/v3/userspage2?page=2&per_page=2"
      },
      "previous": null
   },
   "resources": [
      {
         "guid": "16f43d50-43a2-4981-bae8-633e8248a637",
         "created_at": "2022-08-02T21:37:52Z",
         "updated_at": "2022-08-02T21:37:52Z",
         "username": "smoke_tests",
         "presentation_name": "smoke_tests",
         "origin": "uaa",
         "metadata": {
            "labels": {},
            "annotations": {}
         },
         "links": {
            "self": {
               "href": "https://api.example.org/v3/users/16f43d50-43a2-4981-bae8-633e8248a637"
            }
         }
      },
      {
         "guid": "test1",
         "created_at": "2022-08-02T21:40:34Z",
         "updated_at": "2022-08-02T21:40:34Z",
         "username": "test1",
         "presentation_name": "test1",
         "origin": "uaa",
         "metadata": {
            "labels": {},
            "annotations": {}
         },
         "links": {
            "self": {
               "href": "https://api.example.org/v3/users/test1"
            }
         }
      }
   ]
}`

const listV3UsersPayloadPage2 = `{
   "pagination": {
      "total_results": 3,
      "total_pages": 2,
      "first": {
         "href": "https://api.example.org/v3/users?page=1&per_page=2"
      },
      "last": {
         "href": "https://api.example.org/v3/users?page=2&per_page=2"
      },
      "next": {
         "href": ""
      },
      "previous": {
         "href": "https://api.example.org/v3/users?page=1&per_page=2"
      }
   },
   "resources": [
      {
         "guid": "test2",
         "created_at": "2022-08-02T21:41:59Z",
         "updated_at": "2022-08-02T21:41:59Z",
         "username": "test2",
         "presentation_name": "test2",
         "origin": "uaa",
         "metadata": {
            "labels": {},
            "annotations": {}
         },
         "links": {
            "self": {
               "href": "https://api.example.org/v3/users/test2"
            }
         }
      }
   ]
}`

const listV3ServicePlansPayloadPage1 = `{
    "pagination": {
        "total_results": 2,
        "total_pages": 2,
        "first": {
            "href": "https://api.example.org/v3/service_plans?fields%5Bservice_offering.service_broker%5D=name%2Cguid&include=space.organization%2Cservice_offering&page=1&per_page=1&space_guids=713a9a00-8369-49da-83eb-d34ea1c30545"
        },
        "last": {
            "href": "https://api.example.org/v3/service_plans?fields%5Bservice_offering.service_broker%5D=name%2Cguid&include=space.organization%2Cservice_offering&page=2&per_page=1&space_guids=713a9a00-8369-49da-83eb-d34ea1c30545"
        },
        "next": {
            "href": "https://api.example.org/v3/service_plans?fields%5Bservice_offering.service_broker%5D=name%2Cguid&include=space.organization%2Cservice_offering&page=2&per_page=1&space_guids=713a9a00-8369-49da-83eb-d34ea1c30545"
        },
        "previous": null
    },
    "resources": [
        {
            "guid": "6bf4a2dd-4d0e-4237-bec0-d06242c8802c",
            "created_at": "2019-06-17T13:48:43Z",
            "updated_at": "2022-10-26T09:37:19Z",
            "name": "service-plan1",
            "visibility_type": "public",
            "available": true,
            "free": true,
            "costs": [],
            "description": "ServicePlan 1 Description",
            "maintenance_info": {},
            "broker_catalog": {
                "id": "26c24a23-e6a2-4843-8bc6-f454edb1bd2e",
                "metadata": {},
                "maximum_polling_duration": null,
                "features": {
                    "bindable": true,
                    "plan_updateable": true
                }
            },
            "schemas": {
                "service_instance": {
                    "create": {
                        "parameters": {}
                    },
                    "update": {
                        "parameters": {}
                    }
                },
                "service_binding": {
                    "create": {
                        "parameters": {}
                    }
                }
            },
            "relationships": {
                "service_offering": {
                    "data": {
                        "guid": "bcb04f74-4300-4885-a0be-38e3a53f8e37"
                    }
                }
            },
            "metadata": {
                "labels": {},
                "annotations": {}
            },
            "links": {
                "self": {
                    "href": "https://api.example.org/v3/service_plans/6bf4a2dd-4d0e-4237-bec0-d06242c8802c"
                },
                "service_offering": {
                    "href": "https://api.example.org/v3/service_offerings/bcb04f74-4300-4885-a0be-38e3a53f8e37"
                },
                "visibility": {
                    "href": "https://api.example.org/v3/service_plans/6bf4a2dd-4d0e-4237-bec0-d06242c8802c/visibility"
                }
            }
        }
    ],
    "included": {
        "spaces": [],
        "organizations": [],
        "service_offerings": [
            {
                "guid": "bcb04f74-4300-4885-a0be-38e3a53f8e37",
                "created_at": "2019-01-17T13:48:43Z",
                "updated_at": "2022-11-26T09:37:19Z",
                "name": "service_offering_1",
                "description": "Service Offering 1 Description",
                "available": true,
                "tags": [
                    "tag1"
                ],
                "requires": [],
                "shareable": true,
                "documentation_url": "",
                "broker_catalog": {
                    "id": "7537ffbb-faf3-4efa-9806-af2fc14ca7cb",
                    "metadata": {
                        "displayName": "Display Name 1"
                    },
                    "features": {
                        "plan_updateable": true,
                        "bindable": true,
                        "instances_retrievable": true,
                        "bindings_retrievable": true,
                        "allow_context_updates": true
                    }
                },
                "relationships": {
                    "service_broker": {
                        "data": {
                            "guid": "5314174b-6bbd-4fba-9f5b-a202d19deb21"
                        }
                    }
                },
                "metadata": {
                    "labels": {},
                    "annotations": {}
                },
                "links": {
                    "self": {
                        "href": "https://api.example.org/v3/service_offerings/bcb04f74-4300-4885-a0be-38e3a53f8e37"
                    },
                    "service_plans": {
                        "href": "https://api.example.org/v3/service_plans?service_offering_guids=bcb04f74-4300-4885-a0be-38e3a53f8e37"
                    },
                    "service_broker": {
                        "href": "https://api.example.org/v3/service_brokers/5314174b-6bbd-4fba-9f5b-a202d19deb21"
                    }
                }
            }
        ],
        "service_brokers": [
            {
                "name": "broker-1",
                "guid": "5314174b-6bbd-4fba-9f5b-a202d19deb21"
            }
        ]
    }
}`

const listV3ServicePlansPayloadPage2 = `
{
    "pagination": {
        "total_results": 2,
        "total_pages": 2,
        "first": {
            "href": "https://api.example.org/v3/service_plans?fields%5Bservice_offering.service_broker%5D=name%2Cguid&include=space.organization%2Cservice_offering&page=1&per_page=1&space_guids=713a9a00-8369-49da-83eb-d34ea1c30545"
        },
        "last": {
            "href": "https://api.example.org/v3/service_plans?fields%5Bservice_offering.service_broker%5D=name%2Cguid&include=space.organization%2Cservice_offering&page=2&per_page=1&space_guids=713a9a00-8369-49da-83eb-d34ea1c30545"
        },
        "next": null,
        "previous": {
            "href":  "https://api.example.org/v3/service_plans?fields%5Bservice_offering.service_broker%5D=name%2Cguid&include=space.organization%2Cservice_offering&page=1&per_page=1&space_guids=713a9a00-8369-49da-83eb-d34ea1c30545"
        }
    },
    "resources": [
        {
            "guid": "7bf4a2dd-4d0e-4237-bec0-d06242c8802c",
            "created_at": "2019-06-17T13:48:43Z",
            "updated_at": "2022-10-26T09:37:19Z",
            "name": "service-plan2",
            "visibility_type": "public",
            "available": true,
            "free": false,
            "costs": [],
            "description": "ServicePlan 2 Description",
            "maintenance_info": {},
            "broker_catalog": {
                "id": "36c24a23-e6a2-4843-8bc6-f454edb1bd2e",
                "metadata": {},
                "maximum_polling_duration": null,
                "features": {
                    "bindable": true,
                    "plan_updateable": false
                }
            },
            "schemas": {
                "service_instance": {
                    "create": {
                        "parameters": {}
                    },
                    "update": {
                        "parameters": {}
                    }
                },
                "service_binding": {
                    "create": {
                        "parameters": {}
                    }
                }
            },
            "relationships": {
                "service_offering": {
                    "data": {
                        "guid": "ccb04f74-4300-4885-a0be-38e3a53f8e37"
                    }
                }
            },
            "metadata": {
                "labels": {},
                "annotations": {}
            },
            "links": {
                "self": {
                    "href": "https://api.example.org/v3/service_plans/6bf4a2dd-4d0e-4237-bec0-d06242c8802c"
                },
                "service_offering": {
                    "href": "https://api.example.org/v3/service_offerings/bcb04f74-4300-4885-a0be-38e3a53f8e37"
                },
                "visibility": {
                    "href": "https://api.example.org/v3/service_plans/6bf4a2dd-4d0e-4237-bec0-d06242c8802c/visibility"
                }
            }
        }
    ],
    "included": {
        "spaces": [],
        "organizations": [],
        "service_offerings": [
            {
                "guid": "bcb04f74-4300-4885-a0be-38e3a53f8e37",
                "created_at": "2019-06-17T13:48:43Z",
                "updated_at": "2022-10-26T09:37:19Z",
                "name": "service_offering_1",
                "description": "Service Offering 1 Description",
                "available": true,
                "tags": [
                    "tag1"
                ],
                "requires": [],
                "shareable": true,
                "documentation_url": "",
                "broker_catalog": {
                    "id": "7537ffbb-faf3-4efa-9806-af2fc14ca7cb",
                    "metadata": {
                        "displayName": "Display Name 1"
                    },
                    "features": {
                        "plan_updateable": true,
                        "bindable": true,
                        "instances_retrievable": false,
                        "bindings_retrievable": false,
                        "allow_context_updates": false
                    }
                },
                "relationships": {
                    "service_broker": {
                        "data": {
                            "guid": "5314174b-6bbd-4fba-9f5b-a202d19deb21"
                        }
                    }
                },
                "metadata": {
                    "labels": {},
                    "annotations": {}
                },
                "links": {
                    "self": {
                        "href": "https://api.example.org/v3/service_offerings/bcb04f74-4300-4885-a0be-38e3a53f8e37"
                    },
                    "service_plans": {
                        "href": "https://api.example.org/v3/service_plans?service_offering_guids=bcb04f74-4300-4885-a0be-38e3a53f8e37"
                    },
                    "service_broker": {
                        "href": "https://api.example.org/v3/service_brokers/5314174b-6bbd-4fba-9f5b-a202d19deb21"
                    }
                }
            }
        ],
        "service_brokers": [
            {
                "name": "broker-1",
                "guid": "5314174b-6bbd-4fba-9f5b-a202d19deb21"
            }
        ]
    }
}



`
