<template>
    <div class="row">
        <div class="col-md-10 offset-md-1 py-2">
            <div v-if="this.errorMessage">
                <div class="row">
                    <div class="col-sm-12">
                        <div class="alert alert-danger" role="alert">
                            {{ this.errorMessage }}
                        </div>
                    </div>
                </div>
            </div>
            <div v-if="this.message">
                <div class="row">
                    <div class="col-sm-12">
                        <div class="alert alert-success" role="alert">
                            Endpoint successfully {{ this.message }}
                        </div>
                    </div>
                </div>
            </div>
            <div v-if="!this.errorMessage">
                <div class="row">
                    <div class="col-8 text-left">
                        <h3 v-if="this.loading">Loading...</h3>
                        <h3 v-else>
                            Endpoint: {{ this.endpointName }}
                            <small class="text-muted"></small>
                        </h3>
                    </div>
                    <div v-if="!this.loading" class="col-4 text-right">
                        <router-link :to="{ name: 'collection_detail', params: { id: this.endpoint.CollectionID } }">
                            <button class="btn btn-secondary btn-sm">Back</button>
                        </router-link>
                        &nbsp;
                        <router-link :to="{ name: 'endpoint_edit', params: { id: this.endpoint.ID } }">
                            <button class="btn btn-primary btn-sm">Update</button>
                        </router-link>
                    </div>
                </div>
            </div>
            <div class="row py-2">&nbsp;</div>
            <div v-if="!this.loading">
                <div class="row">
                    <div class="col-sm-12">
                        <table class="table table-bordered">
                            <tr>
                                <th>Name</th>
                                <td>{{ this.endpoint.Name }}</td>
                            </tr>
                            <tr>
                                <th>Method</th>
                                <td>{{ this.computedMethod }}</td>
                            </tr>
                            <tr>
                                <th>Expected Repsonse</th>
                                <td>{{ this.endpoint.ResponseStatusCode }}</td>
                            </tr>
                            <tr>
                                <th>Path</th>
                                <td>{{ this.endpoint.Path }}</td>
                            </tr>
                            <tr>
                                <td colspan="2">
                                    <strong>Output</strong>
                                    <hr />
                                    <span class="json-container" v-html="this.prettyOutput"></span>
                                </td>
                            </tr>
                        </table>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import axios from "axios";
import { prettyPrintJson } from "pretty-print-json"; // https://www.npmjs.com/package/pretty-print-json

export default {
    name: "EndpointDetailPage",

    props: {
        "message": String, 
    },

    data() {
        return {
            loading: false,
            endpoint: null,
            endpointName: "",
            errorMessage: "",
        }
    },

    created() {
        this.$watch(
            () => this.$route.params, () => {
                this.fetchData();
            }, { immediate: true }
        );
    },

    computed: {
        computedMethod: function() {
            if (this.endpoint == null) {
                return "";
            }
            if (this.endpoint.Method == "") {
                return "Any";
            }
            return this.endpoint.Method;
        },
        prettyOutput: function() {
            if (this.endpoint == null || this.endpoint.Collection == null) {
                return "";
            }
            // TODO: make this work with _real_ content types
            if (this.endpoint.Collection.ContentType == 'JSON') {
                try {
                    return prettyPrintJson.toHtml(JSON.parse(this.endpoint.Output), {
                        lineNumbers: true,
                    });
                } catch (e) {
                    // Assume JSON syntax error
                    return this.endpoint.Output;
                }
            }

            return this.endpoint.Output;
        }
    },

    methods: {
        fetchData() {
            this.errorMessage = "";
            this.endpoint = null;
            this.endpointName = "";
            this.loading = true;

            axios.get("http://localhost:3000/itinerary/endpoint/"+this.$route.params.id)
                .then((response) => {
                    this.endpoint = response.data;
                    this.endpointName = this.endpoint.Name;
                    if (!this.endpoint.Name) {
                        this.endpointName = "Endpoint";
                    }
                    this.loading = false;
                })
                .catch((error) => {
                    this.errorMessage = error;
                })
        }   
    }
}
</script>