<signals class="ui container">
  <h2 class="ui header">{ name }</h2>
  <div class="ui segments">
    <div each={ signals } class="ui segment">
      <span class="signal">{ remote }:{ name }</span>
      <button onclick={ send } class="ui right floated labeled icon button">
        <i class="play icon"></i> Send
      </button>
    </div>
  </div>

  <script>
    this.name = "Signals";
    this.signals = opts.api.signals;
    send(event) {
      console.log(event);
      opts.send(event.item);
    }
  </script>

  <style>
    .signal {
      font-size: 1.6rem;
      line-height: 36px;
    }
  </style>
</signals>
