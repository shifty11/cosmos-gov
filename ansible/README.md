# Node Exporter, Promtail and Cosmos Gov

Sets up everything to run and monitor Cosmos Gov Telegram Bot

## Step 1: Set up your own inventory file

Copy inventory file, and make your edits.

```bash
cp samples/inventory.sample inventory
```

##  Step 2: Run main playbook to set up a fresh monitor

The main monitor ansible file is main.yml, which sets up a new fresh monitor from scratch. It will set up firewall, install Prometheus, Grafana and Alert Manager.

```bash
ansible-playbook -i inventory main.yml
```

## Step 4: Run separate playbooks

You might want to run separate playbook as needed:

| Playbook                   | Description                |
|----------------------------|----------------------------|
| main.yml                   | Full Setup                 |
| main_update.yml            | Updates Cosmos Gov service |

That's it!
