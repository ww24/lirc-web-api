<signals class="ui container">
  <h2 class="ui header">{ name }</h2>
  <div class="ui segments">
    <div each={ signals } class="ui segment">
      <span class="signal">{ remote }:{ name }</span>
      <button onclick={ send } ref="send_btn" class="ui right floated labeled icon button">
        <i class="play icon"></i> Send
      </button>
    </div>
  </div>

  <script>
    this.name = "Signals";
    this.signals = opts.api.signals;
    send(event) {
      let item = event.item;
      let index = this.signals.indexOf(item);
      let $button = $(this.refs.send_btn[index]);
      $button
        .addClass("loading")
        .prop("disabled", true);
      opts.send(item).then(() => {
        $button
          .removeClass("loading")
          .prop("disabled", false);
      });
    }
  </script>

  <style>
    .signal {
      font-size: 1.6rem;
      line-height: 36px;
    }
  </style>
</signals>
