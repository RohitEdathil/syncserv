# TypeSync

A simple web app to sync code across multiple devices.

The idea is that there is a broadcaster who types in code and listeners who see the code in real time. Backend is written in **Go** and frontend is written in **Svelte**. This could be used for pair programming, teaching, or just for fun.

This repo contains the backend code.

## Features

- Broadcast code from a broadcaster to multiple listeners
- Broadcaster can see how many listeners are connected
- A red-flag, green-flag system to let the broadcaster know if the listeners are able to follow along
