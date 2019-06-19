`context` helps us to manage long-running processes.

The context package provides functions to derive new Context values from existing ones. These values form a tree: when 
a Context is canceled, all Contexts derived from it are also canceled.