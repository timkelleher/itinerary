<template>
    <div class="row">
        <div class="col-md-10 offset-md-1 py-2">
            <div>
                <div class="row">
                    <div class="col-8 text-left">
                        <h3>New Endpoint</h3>
                    </div>
                    <div v-if="this.loaded" class="col-4 text-right">
                        <router-link :to="{ name: 'collection_detail', params: { id: this.$route.params.collectionID } }">
                            <button class="btn btn-secondary btn-sm">Back</button>
                        </router-link>
                        &nbsp;
                        <button @click="persistEndpoint" role="button" class="btn btn-primary btn-sm">Save</button>
                    </div>
                </div>
                <div class="row" v-if="this.errorMessage">
                    <div class="col-sm-12">
                        <div class="alert alert-danger" role="alert">
                            {{ this.errorMessage }}
                        </div>
                    </div>
                </div>
                <div class="row" v-if="this.successMessage">
                    <div class="col-sm-12">
                        <div class="alert alert-success" role="alert">
                            {{ this.successMessage }}
                        </div>
                    </div>
                </div>
                <div v-if="this.loaded">
                    <form @submit="persistEndpoint">
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
    </div>
</template>

<script>
import axios from "axios";
import MethodsDropdown from "../components/MethodsDropdown.vue";
import ResponseStatusCodeDropdown from "../components/ResponseStatusCodeDropdown.vue";

export default {
    name: "EndpointAddPage",

    components: {
        MethodsDropdown,
        ResponseStatusCodeDropdown,
    },

    data() {
        return {
            loaded: {
                "Methods": false,
                "StatusCodes": false,
                "Collection": false,
            },
            collectionID: 0,
            methods: [],
            responseStatusCodes: [],
            errorMessage: "",
            successMessage: "",
            cachedPathPrefix: "/",
            endpoint: {
                Name: "",
                Method: "",
                ResponseStatusCode: "200",
                Path: "",
                Output: ""
            }
        }
    },

    mounted() {
        this.collectionID = this.$route.params.collectionID;
        this.fetchMethods();
        this.fetchResponseStatusCodes();
        this.fetchCollection();
    },

/*
maybe let users "save" unfinished work
    created() {
        this.$watch(
            () => this.$route.params, () => {
                this.fetchData();
            }, { immediate: true }
        );
    },
*/

    computed: {
        loading: function() {
            return this.loaded.Methods &&
                this.loaded.StatusCodes &&
                this.loaded.Collection &&
                !this.errorMessage;
        }
    },

    methods: {
        fetchMethods() {
            this.methods = [];

            axios.get("http://localhost:3000/itinerary/methods")
                .then((response) => {
                    this.methods = response.data;
                    this.loaded.Methods = true;
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
                    this.loaded.StatusCodes = true;
                })
                .catch((error) => {
                    this.errorMessage = error;
                })
        },
        fetchCollection() {
            axios.get("http://localhost:3000/itinerary/collection/"+this.collectionID)
                .then((response) => {
                    this.cachedPathPrefix = response.data.PathPrefix;
                    this.loaded.Collection = true;
                })
                .catch((error) => {
                    this.errorMessage = error;
                })
        },
        persistEndpoint() {
            axios.put("http://localhost:3000/itinerary/endpoint", {
                "name": this.endpoint.Name,
                "collection_id": parseInt(this.collectionID),
                "method": this.endpoint.Method,
                "response_status_code": this.endpoint.ResponseStatusCode,
                "path": this.endpoint.Path,
                "output": this.endpoint.Output,
            })
                .then((response) => {
                    this.endpoint = response.data;
                    this.successMessage = "Endpoint has been created";
                    this.$router.push({
                        name: "collection_detail",
                        params: {"id": this.endpoint.CollectionID},
                        props: {"mode": "created"},
                    })
                })
                .catch((error) => {
                    this.errorMessage = error;
                })
        }
    }
}
</script>