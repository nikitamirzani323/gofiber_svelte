<script>
    import { mcompany,fetchCompany } from "../../stores/companystore"
    import Panel from '../../componets/Panel.svelte'

	let page = "Company"
	let screen_height = screen.height - 480;
    let totalrow = 200
    let totalrecord = 200
    let paging = [
        {idpaging:0,nmpaging:"1 Of "+totalrow},
        {idpaging:1,nmpaging:"2 Of "+totalrow},
        {idpaging:2,nmpaging:"3 Of "+totalrow},
        {idpaging:3,nmpaging:"4 Of "+totalrow},
    ]
    fetchCompany(`http://localhost:4000/api/company`)
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
                {#each paging as paging }
                    <option selected>{paging.nmpaging}</option>
                {/each}
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
					<th width="5%" style="text-align: center;font-size: 15px;">CURR</th>
					<th width="20%" style="text-align: left;font-size: 15px;">COMPANY</th>
					<th width="*" style="text-align: left;font-size: 15px;">OWNER</th>
				</tr>
			</thead>
			<tbody>
                {#each $mcompany as rec }
                    <tr>
                        <td style="text-align: center;font-size: 14px;">{rec.no}</td>
                        <td style="text-align: center;font-size: 14px;">{rec.idcurr}</td>
                        <td style="text-align: left;font-size: 14px;">{rec.idcompany}</td>
                        <td style="text-align: left;font-size: 14px;">{rec.nmowner}</td>
                    </tr>
                {/each}
			</tbody>
		</table>
	</slot:template>
	<slot:template slot="cfooter">
		Total Record : {totalrecord}
	</slot:template>
</Panel>