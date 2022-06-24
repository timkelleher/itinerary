<template>
    <div class="col-md-12 py-2">
        <div class="row">
            <div class="col-8">
                <h3>Endpoints</h3>
            </div>
            <div class="col-4 text-right">
                <router-link :to="{ name: 'endpoint_add', params: { collectionID: this.collectionID } }">
                    <button role="button" class="btn btn-success btn-sm">New Endpoint</button>
                </router-link>
            </div>
        </div>
        <div class="row">
            <div class="col-sm-12">
                <div v-if="this.endpoints != null && this.endpoints.length > 0">
                    <table class="table table-bordered">
                        <tr>
                            <th>Name</th>
                            <th>Path</th>
                            <th>Method</th>
                            <th>Expected Response</th>
                            <th>Inputs</th>
                            <th>&nbsp;</th>
                        </tr>
                        <tr v-bind:key="endpoint.ID" v-for="endpoint in endpoints">
                            <td>{{ endpoint.Name }}</td>
                            <td>{{ endpoint.Path }}</td>
                            <td>{{ this.calcMethod(endpoint.Method) }}</td>
                            <td>{{ endpoint.ResponseStatusCode }}</td>
                            <td></td>
                            <td>
                                <router-link :to="{ name: 'endpoint_detail', params: { id: endpoint.ID } }" >
                                    <button type="button" class="btn btn-primary btn-sm">View</button>
                                </router-link>
                                &nbsp;
                                <!--
                                <button type="button" class="btn btn-warning btn-sm">Run</button>
                                -->
                            </td>
                        </tr>
                    </table>
                </div>
                <div v-else class="col-sm-12 align-center">
                    <div class="alert alert-warning">No endpoints exist yet.</div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: "EndpointsTable",

    props: {
        "endpoints": Array,
        "collectionID": Number,
    },

    methods: {
        // TODO: remove this duplicate method
        calcMethod(method) {
            if (method == "") {
                return "Any";
            }
            return method;
        }
    }
}
</script>