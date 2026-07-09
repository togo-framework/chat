# chat — docs

Floating chat widget + a responder registry. The base ships a DB keyword responder; a chat-ai driver swaps in real model answers via the same contract.

- **Source:** https://github.com/togo-framework/chat
- **Reference app:** https://fadymondy.com

## Install

```bash
togo install togo-framework/chat
```

The plugin self-registers on blank-import (no manual wiring). See the top-level
README for usage. This plugin was extracted from the fadymondy.com togo app as
part of turning that site into a reusable plugin suite.
