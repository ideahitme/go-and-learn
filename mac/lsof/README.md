### MAC OS X commands with lsof 

#### Show established connection with PID 

Lists all established connections launched by `main.go` (refreshed every second)

```bash
watch -n 1 sh -c "lsof -i | grep ESTABLISHED | grep main"
```

