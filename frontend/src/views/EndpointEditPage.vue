<template>
    <div class="row">
        <div class="col-md-10 offset-md-1 py-2">
            <div class="row">
                <div class="col-8 text-left">
                    <h3 v-if="this.loading">Loading...</h3>
                    <h3 v-else>
                        Update {{ this.cachedName }}
                        <small class="text-muted"></small>
                    </h3>
                </div>
                <div v-if="!this.loading" class="col-4 text-right">
                    <router-link :to="{ name: 'endpoint_detail', params: { id: this.endpoint.ID } }">
                        <button class="btn btn-secondary btn-sm">Back</button>
                    </router-link>
                    &nbsp;
                    <button @click="persistData" role="button" class="btn btn-primary btn-sm">Save</button>
                </div>
            </div>
            <AlertsSection :errorMessage="this.errorMessage" :successMessage="this.successMessage" />
            <div v-if="!this.loading">
                <form @submit="persistData">
                    <div class="row">
                        <div class="col-md-6">
                            <div class="form-group">
                                <label for="name">Name</label>
                                <input type="name" class="form-control" id="name" placeholder="Endpoint name" v-model="endpoint.Name" />
                            </div>
                        </div>
                        <div class="col-md-6">
                            <div class="form-group">
                                <label for="name">Method</label>
                                <MethodsDropdown
                                    :options="this.methods"
                                    :selected="this.endpoint.Method"
                                    v-on:updateInput="this.endpoint.Method = $event"
                                />
                            </div>
                        </div>
                    </div>
                    <div class="row">
                        <div class="col-md-6">
                            <div class="form-group">
                                <label for="name">Expected Response</label>
                                <ResponseStatusCodeDropdown
                                    :options="this.responseStatusCodes"
                                    :selected="this.endpoint.ResponseStatusCode"
                                    v-on:updateInput="this.endpoint.ResponseStatusCode = $event"
                                />
                            </div>
                        </div>
                        <div class="col-md-6">
                            <div class="form-group">
                                <label for="name">Path</label>
                                <div class="row">
                                    <div class="col-sm-4 col-md-3">
                                        {{ this.cachedPathPrefix }}
                                    </div>
                                    <div class="col-sm-8 col-md-9">
                                        <input type="name" class="form-control" id="path" placeholder="Path" v-model="endpoint.Path" />
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="row">
                        <div class="col-md-12">
                        <div class="form-group">
                            <label for="output">Output</label>
                            <textarea class="form-control" id="output" rows="12" v-model="endpoint.Output" />
                        </div>
                        </div>
                    </div>
                </form>
            </div>
        </div>
    </div>
</template>

<script>
import axios from "axios";
import AlertsSection from "../components/AlertsSection.vue";
import MethodsDropdown from "../components/MethodsDropdown.vue";
import ResponseStatusCodeDropdown from "../components/ResponseStatusCodeDropdown.vue";

export default {
    name: "EndpointEditPage",

    components: {
        MethodsDropdown,
        ResponseStatusCodeDropdown,
        AlertsSection,
    },

    data() {
        return {
            loading: false,
            endpoint: null,
            cachedName: "",
            cachedPathPrefix: "",
            methods: [],
            responseStatusCodes: [],
            errorMessage: "",
            successMessage: "",
        }
    },

    created() {
        this.$watch(
            () => this.$route.params, () => {
                this.fetchData();
            }, { immediate: true }
        );
    },
    
    methods: {
        fetchData() {
            this.error = null;
            this.endpoint = null;
            this.loading = true;
            this.cachedName = "";
            this.cachedPathPrefix = "";

            if (this.$route.params.id) {
                axios.get("http://localhost:3000/itinerary/endpoint/"+this.$route.params.id)
                    .then((response) => {
                        this.endpoint = response.data;
                        this.loading = false;
                        this.cachedName = this.endpoint.Name || "Endpoint";
                        // TODO: grab from env var or something
                        // TODO: normalize the path forward-slashes
                        this.cachedPathPrefix = "/api/" + this.endpoint.PathPrefix + "/";
                        this.fetchMethods();
                        this.fetchResponseStatusCodes();
                    })
                    .catch((error) => {
                        this.errorMessage = error;
                    })
            } else {
                this.loading = false;
            }
        },
        fetchMethods() {
            this.methods = [];

            axios.get("http://localhost:3000/itinerary/methods")
                .then((response) => {
                    this.methods = response.data;
                })
                .catch((error) => {
                    this.errorMessage = error;
                })
        },
        fetchResponseStatusCodes() {
            this.responseStatusCodes = [];

            axios.get("http://localhost:3000/itinerary/statuses")
                .then((response) => {
                    this.responseStatusCodes = response.data;
                })
                .catch((error) => {
                    this.errorMessage = error;
                })
        },
        persistData() {
            axios.post("http://localhost:3000/itinerary/endpoint/"+this.endpoint.ID, {
                "id": this.endpoint.ID,
                "name": this.endpoint.Name,
                "method": this.endpoint.Method,
                "response_status_code": this.endpoint.ResponseStatusCode,
                "path": this.endpoint.Path,
                "output": this.endpoint.Output,
            })
                .then((response) => {
                    this.endpoint = response.data;
                    this.successMessage = "Endpoint has been updated";
                    this.$router.push({
                        name: "endpoint_detail",
                        params: {"id": this.endpoint.ID},
                        props: {"mode": "updated"},
                    })
                })
                .catch((error) => {
                    this.errorMessage = error;
                })
        }
    }
}
</script>