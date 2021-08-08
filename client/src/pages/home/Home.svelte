<script>
    import Panel from '../../componets/Panel.svelte'

    const queryString = window.location.search
	const urlParams = new URLSearchParams(queryString)
	const token_browser = urlParams.get('token')
	let page = "Home"
	let screen_height = screen.height - 480;
</script>
<button class="btn btn-primary" style="border-radius: 0px;">
    New
</button>
<button class="btn btn-primary" style="border-radius: 0px;">
    Delete
</button>
<button class="btn btn-primary" style="border-radius: 0px;">
    Refresh
</button>
<Panel height_body="{screen_height}px">
	<slot:template slot="cheader">
		{page}
		<div class="float-end">
			<select class="form-select" aria-label="Default select example">
				<option selected>1 of 100</option>
				<option value="1">One</option>
				<option value="2">Two</option>
				<option value="3">Three</option>
			</select>
		</div>
	</slot:template>
	<slot:template slot="csearch">
		<div class="col-lg-12" style="padding: 5px;">
			<div class="row g-2">
				<div class="col-sm-11">
					<input
					type="text"
					class="form-control"
					placeholder="Search"
					aria-label="Search" />
				</div>
				<div class="col-sm">
					<button class="btn btn-primary" style="border-radius: 0px;">
						Search
					</button>
				</div>
			</div>
		</div>
	</slot:template>
	<slot:template slot="cbody">
		<table class="table table-striped table-hover">
			<thead>
				<tr>
					<th width="1%" style="text-align: center;font-size: 15px;">NO</th>
					<th width="20%" style="text-align: left;font-size: 15px;">NAME</th>
					<th width="*" style="text-align: left;font-size: 15px;">DESCRIPTION</th>
				</tr>
			</thead>
			<tbody>
				<tr>
					<td style="text-align: center;font-size: 14px;">1</td>
					<td style="text-align: left;font-size: 14px;">Budi</td>
					<td style="text-align: left;font-size: 14px;">JL kembang no 8</td>
				</tr>
				<tr>
					<td style="text-align: center;font-size: 14px;">1</td>
					<td style="text-align: left;font-size: 14px;">Budi</td>
					<td style="text-align: left;font-size: 14px;">JL kembang no 8</td>
				</tr>
				<tr>
					<td style="text-align: center;font-size: 14px;">1</td>
					<td style="text-align: left;font-size: 14px;">Budi</td>
					<td style="text-align: left;font-size: 14px;">JL kembang no 8</td>
				</tr>
				<tr>
					<td style="text-align: center;font-size: 14px;">1</td>
					<td style="text-align: left;font-size: 14px;">Budi</td>
					<td style="text-align: left;font-size: 14px;">JL kembang no 8</td>
				</tr>
				<tr>
					<td style="text-align: center;font-size: 14px;">1</td>
					<td style="text-align: left;font-size: 14px;">Budi</td>
					<td style="text-align: left;font-size: 14px;">JL kembang no 8</td>
				</tr>
			</tbody>
		</table>
	</slot:template>
	<slot:template slot="cfooter">
		Total Record : 0
	</slot:template>
</Panel>