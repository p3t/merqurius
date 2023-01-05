
![merQurius logo](img/merqurius.svg "merQurius")

# merQurius - a MQTT message proxy-buffer
Simple side-kick service, meant to run on an edge device (like a Raspberry Pi) to proxy and buffer MQTT messages.
merQurius can be configured in the startup sequence to be available _before_ the data-generating service comes up.
Also, merQurius will buffer messages in case that the central MQTT broker is not available - and re-deliver them after‚ reconnect.

