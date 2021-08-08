<script>
    export let height_body = "650px"
</script>

<div class="card" style="border-radius: 0px;margin-top:10px;">
  <div class="card-header" style="">
    <slot name="cheader"></slot>
</div>
  <slot name="csearch"></slot>

  <div class="card-body" style="height:{height_body};overflow:scroll;">
    <slot name="cbody"></slot>
  </div>
  <div class="card-footer" style="">
    <slot name="cfooter"></slot>
  </div>
</div>
