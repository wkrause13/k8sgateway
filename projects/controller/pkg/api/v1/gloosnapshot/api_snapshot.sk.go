// Code generated by solo-kit. DO NOT EDIT.

package gloosnapshot

import (
	"fmt"
	"hash"
	"hash/fnv"
	"log"

	github_com_solo_io_gloo_projects_controller_pkg_api_external_solo_ratelimit "github.com/solo-io/gloo/projects/controller/pkg/api/external/solo/ratelimit"
	gloo_solo_io "github.com/solo-io/gloo/projects/controller/pkg/api/v1"
	enterprise_gloo_solo_io "github.com/solo-io/gloo/projects/controller/pkg/api/v1/enterprise/options/extauth/v1"
	graphql_gloo_solo_io "github.com/solo-io/gloo/projects/controller/pkg/api/v1/enterprise/options/graphql/v1beta1"
	gateway_solo_io "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"

	"github.com/rotisserie/eris"
	"github.com/solo-io/go-utils/hashutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type ApiSnapshot struct {
	Artifacts          gloo_solo_io.ArtifactList
	Endpoints          gloo_solo_io.EndpointList
	Proxies            gloo_solo_io.ProxyList
	UpstreamGroups     gloo_solo_io.UpstreamGroupList
	Secrets            gloo_solo_io.SecretList
	Upstreams          gloo_solo_io.UpstreamList
	AuthConfigs        enterprise_gloo_solo_io.AuthConfigList
	Ratelimitconfigs   github_com_solo_io_gloo_projects_controller_pkg_api_external_solo_ratelimit.RateLimitConfigList
	VirtualServices    gateway_solo_io.VirtualServiceList
	RouteTables        gateway_solo_io.RouteTableList
	Gateways           gateway_solo_io.GatewayList
	VirtualHostOptions gateway_solo_io.VirtualHostOptionList
	RouteOptions       gateway_solo_io.RouteOptionList
	HttpGateways       gateway_solo_io.MatchableHttpGatewayList
	TcpGateways        gateway_solo_io.MatchableTcpGatewayList
	GraphqlApis        graphql_gloo_solo_io.GraphQLApiList
}

func (s ApiSnapshot) Clone() ApiSnapshot {
	return ApiSnapshot{
		Artifacts:          s.Artifacts.Clone(),
		Endpoints:          s.Endpoints.Clone(),
		Proxies:            s.Proxies.Clone(),
		UpstreamGroups:     s.UpstreamGroups.Clone(),
		Secrets:            s.Secrets.Clone(),
		Upstreams:          s.Upstreams.Clone(),
		AuthConfigs:        s.AuthConfigs.Clone(),
		Ratelimitconfigs:   s.Ratelimitconfigs.Clone(),
		VirtualServices:    s.VirtualServices.Clone(),
		RouteTables:        s.RouteTables.Clone(),
		Gateways:           s.Gateways.Clone(),
		VirtualHostOptions: s.VirtualHostOptions.Clone(),
		RouteOptions:       s.RouteOptions.Clone(),
		HttpGateways:       s.HttpGateways.Clone(),
		TcpGateways:        s.TcpGateways.Clone(),
		GraphqlApis:        s.GraphqlApis.Clone(),
	}
}

func (s ApiSnapshot) Hash(hasher hash.Hash64) (uint64, error) {
	if hasher == nil {
		hasher = fnv.New64()
	}
	if _, err := s.hashArtifacts(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashEndpoints(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashProxies(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashUpstreamGroups(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashSecrets(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashUpstreams(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashAuthConfigs(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashRatelimitconfigs(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashVirtualServices(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashRouteTables(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashGateways(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashVirtualHostOptions(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashRouteOptions(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashHttpGateways(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashTcpGateways(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashGraphqlApis(hasher); err != nil {
		return 0, err
	}
	return hasher.Sum64(), nil
}

func (s ApiSnapshot) hashArtifacts(hasher hash.Hash64) (uint64, error) {
	clonedList := s.Artifacts.Clone()
	for _, v := range clonedList {
		v.Metadata.Annotations = nil
	}
	return hashutils.HashAllSafe(hasher, clonedList.AsInterfaces()...)
}

func (s ApiSnapshot) hashEndpoints(hasher hash.Hash64) (uint64, error) {
	clonedList := s.Endpoints.Clone()
	for _, v := range clonedList {
		v.Metadata.Annotations = nil
	}
	return hashutils.HashAllSafe(hasher, clonedList.AsInterfaces()...)
}

func (s ApiSnapshot) hashProxies(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.Proxies.AsInterfaces()...)
}

func (s ApiSnapshot) hashUpstreamGroups(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.UpstreamGroups.AsInterfaces()...)
}

func (s ApiSnapshot) hashSecrets(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.Secrets.AsInterfaces()...)
}

func (s ApiSnapshot) hashUpstreams(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.Upstreams.AsInterfaces()...)
}

func (s ApiSnapshot) hashAuthConfigs(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.AuthConfigs.AsInterfaces()...)
}

func (s ApiSnapshot) hashRatelimitconfigs(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.Ratelimitconfigs.AsInterfaces()...)
}

func (s ApiSnapshot) hashVirtualServices(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.VirtualServices.AsInterfaces()...)
}

func (s ApiSnapshot) hashRouteTables(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.RouteTables.AsInterfaces()...)
}

func (s ApiSnapshot) hashGateways(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.Gateways.AsInterfaces()...)
}

func (s ApiSnapshot) hashVirtualHostOptions(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.VirtualHostOptions.AsInterfaces()...)
}

func (s ApiSnapshot) hashRouteOptions(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.RouteOptions.AsInterfaces()...)
}

func (s ApiSnapshot) hashHttpGateways(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.HttpGateways.AsInterfaces()...)
}

func (s ApiSnapshot) hashTcpGateways(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.TcpGateways.AsInterfaces()...)
}

func (s ApiSnapshot) hashGraphqlApis(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.GraphqlApis.AsInterfaces()...)
}

func (s ApiSnapshot) HashFields() []zap.Field {
	var fields []zap.Field
	hasher := fnv.New64()
	ArtifactsHash, err := s.hashArtifacts(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("artifacts", ArtifactsHash))
	EndpointsHash, err := s.hashEndpoints(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("endpoints", EndpointsHash))
	ProxiesHash, err := s.hashProxies(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("proxies", ProxiesHash))
	UpstreamGroupsHash, err := s.hashUpstreamGroups(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("upstreamGroups", UpstreamGroupsHash))
	SecretsHash, err := s.hashSecrets(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("secrets", SecretsHash))
	UpstreamsHash, err := s.hashUpstreams(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("upstreams", UpstreamsHash))
	AuthConfigsHash, err := s.hashAuthConfigs(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("authConfigs", AuthConfigsHash))
	RatelimitconfigsHash, err := s.hashRatelimitconfigs(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("ratelimitconfigs", RatelimitconfigsHash))
	VirtualServicesHash, err := s.hashVirtualServices(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("virtualServices", VirtualServicesHash))
	RouteTablesHash, err := s.hashRouteTables(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("routeTables", RouteTablesHash))
	GatewaysHash, err := s.hashGateways(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("gateways", GatewaysHash))
	VirtualHostOptionsHash, err := s.hashVirtualHostOptions(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("virtualHostOptions", VirtualHostOptionsHash))
	RouteOptionsHash, err := s.hashRouteOptions(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("routeOptions", RouteOptionsHash))
	HttpGatewaysHash, err := s.hashHttpGateways(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("httpGateways", HttpGatewaysHash))
	TcpGatewaysHash, err := s.hashTcpGateways(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("tcpGateways", TcpGatewaysHash))
	GraphqlApisHash, err := s.hashGraphqlApis(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("graphqlApis", GraphqlApisHash))
	snapshotHash, err := s.Hash(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	return append(fields, zap.Uint64("snapshotHash", snapshotHash))
}

func (s *ApiSnapshot) GetResourcesList(resource resources.Resource) (resources.ResourceList, error) {
	switch resource.(type) {
	case *gloo_solo_io.Artifact:
		return s.Artifacts.AsResources(), nil
	case *gloo_solo_io.Endpoint:
		return s.Endpoints.AsResources(), nil
	case *gloo_solo_io.Proxy:
		return s.Proxies.AsResources(), nil
	case *gloo_solo_io.UpstreamGroup:
		return s.UpstreamGroups.AsResources(), nil
	case *gloo_solo_io.Secret:
		return s.Secrets.AsResources(), nil
	case *gloo_solo_io.Upstream:
		return s.Upstreams.AsResources(), nil
	case *enterprise_gloo_solo_io.AuthConfig:
		return s.AuthConfigs.AsResources(), nil
	case *github_com_solo_io_gloo_projects_controller_pkg_api_external_solo_ratelimit.RateLimitConfig:
		return s.Ratelimitconfigs.AsResources(), nil
	case *gateway_solo_io.VirtualService:
		return s.VirtualServices.AsResources(), nil
	case *gateway_solo_io.RouteTable:
		return s.RouteTables.AsResources(), nil
	case *gateway_solo_io.Gateway:
		return s.Gateways.AsResources(), nil
	case *gateway_solo_io.VirtualHostOption:
		return s.VirtualHostOptions.AsResources(), nil
	case *gateway_solo_io.RouteOption:
		return s.RouteOptions.AsResources(), nil
	case *gateway_solo_io.MatchableHttpGateway:
		return s.HttpGateways.AsResources(), nil
	case *gateway_solo_io.MatchableTcpGateway:
		return s.TcpGateways.AsResources(), nil
	case *graphql_gloo_solo_io.GraphQLApi:
		return s.GraphqlApis.AsResources(), nil
	default:
		return resources.ResourceList{}, eris.New("did not contain the input resource type returning empty list")
	}
}

func (s *ApiSnapshot) RemoveFromResourceList(resource resources.Resource) error {
	refKey := resource.GetMetadata().Ref().Key()
	switch resource.(type) {
	case *gloo_solo_io.Artifact:

		for i, res := range s.Artifacts {
			if refKey == res.GetMetadata().Ref().Key() {
				s.Artifacts = append(s.Artifacts[:i], s.Artifacts[i+1:]...)
				break
			}
		}
		return nil
	case *gloo_solo_io.Endpoint:

		for i, res := range s.Endpoints {
			if refKey == res.GetMetadata().Ref().Key() {
				s.Endpoints = append(s.Endpoints[:i], s.Endpoints[i+1:]...)
				break
			}
		}
		return nil
	case *gloo_solo_io.Proxy:

		for i, res := range s.Proxies {
			if refKey == res.GetMetadata().Ref().Key() {
				s.Proxies = append(s.Proxies[:i], s.Proxies[i+1:]...)
				break
			}
		}
		return nil
	case *gloo_solo_io.UpstreamGroup:

		for i, res := range s.UpstreamGroups {
			if refKey == res.GetMetadata().Ref().Key() {
				s.UpstreamGroups = append(s.UpstreamGroups[:i], s.UpstreamGroups[i+1:]...)
				break
			}
		}
		return nil
	case *gloo_solo_io.Secret:

		for i, res := range s.Secrets {
			if refKey == res.GetMetadata().Ref().Key() {
				s.Secrets = append(s.Secrets[:i], s.Secrets[i+1:]...)
				break
			}
		}
		return nil
	case *gloo_solo_io.Upstream:

		for i, res := range s.Upstreams {
			if refKey == res.GetMetadata().Ref().Key() {
				s.Upstreams = append(s.Upstreams[:i], s.Upstreams[i+1:]...)
				break
			}
		}
		return nil
	case *enterprise_gloo_solo_io.AuthConfig:

		for i, res := range s.AuthConfigs {
			if refKey == res.GetMetadata().Ref().Key() {
				s.AuthConfigs = append(s.AuthConfigs[:i], s.AuthConfigs[i+1:]...)
				break
			}
		}
		return nil
	case *github_com_solo_io_gloo_projects_controller_pkg_api_external_solo_ratelimit.RateLimitConfig:

		for i, res := range s.Ratelimitconfigs {
			if refKey == res.GetMetadata().Ref().Key() {
				s.Ratelimitconfigs = append(s.Ratelimitconfigs[:i], s.Ratelimitconfigs[i+1:]...)
				break
			}
		}
		return nil
	case *gateway_solo_io.VirtualService:

		for i, res := range s.VirtualServices {
			if refKey == res.GetMetadata().Ref().Key() {
				s.VirtualServices = append(s.VirtualServices[:i], s.VirtualServices[i+1:]...)
				break
			}
		}
		return nil
	case *gateway_solo_io.RouteTable:

		for i, res := range s.RouteTables {
			if refKey == res.GetMetadata().Ref().Key() {
				s.RouteTables = append(s.RouteTables[:i], s.RouteTables[i+1:]...)
				break
			}
		}
		return nil
	case *gateway_solo_io.Gateway:

		for i, res := range s.Gateways {
			if refKey == res.GetMetadata().Ref().Key() {
				s.Gateways = append(s.Gateways[:i], s.Gateways[i+1:]...)
				break
			}
		}
		return nil
	case *gateway_solo_io.VirtualHostOption:

		for i, res := range s.VirtualHostOptions {
			if refKey == res.GetMetadata().Ref().Key() {
				s.VirtualHostOptions = append(s.VirtualHostOptions[:i], s.VirtualHostOptions[i+1:]...)
				break
			}
		}
		return nil
	case *gateway_solo_io.RouteOption:

		for i, res := range s.RouteOptions {
			if refKey == res.GetMetadata().Ref().Key() {
				s.RouteOptions = append(s.RouteOptions[:i], s.RouteOptions[i+1:]...)
				break
			}
		}
		return nil
	case *gateway_solo_io.MatchableHttpGateway:

		for i, res := range s.HttpGateways {
			if refKey == res.GetMetadata().Ref().Key() {
				s.HttpGateways = append(s.HttpGateways[:i], s.HttpGateways[i+1:]...)
				break
			}
		}
		return nil
	case *gateway_solo_io.MatchableTcpGateway:

		for i, res := range s.TcpGateways {
			if refKey == res.GetMetadata().Ref().Key() {
				s.TcpGateways = append(s.TcpGateways[:i], s.TcpGateways[i+1:]...)
				break
			}
		}
		return nil
	case *graphql_gloo_solo_io.GraphQLApi:

		for i, res := range s.GraphqlApis {
			if refKey == res.GetMetadata().Ref().Key() {
				s.GraphqlApis = append(s.GraphqlApis[:i], s.GraphqlApis[i+1:]...)
				break
			}
		}
		return nil
	default:
		return eris.Errorf("did not remove the resource because its type does not exist [%T]", resource)
	}
}

func (s *ApiSnapshot) RemoveMatches(predicate core.Predicate) {
	var Artifacts gloo_solo_io.ArtifactList
	for _, res := range s.Artifacts {
		if matches := predicate(res.GetMetadata()); !matches {
			Artifacts = append(Artifacts, res)
		}
	}
	s.Artifacts = Artifacts
	var Endpoints gloo_solo_io.EndpointList
	for _, res := range s.Endpoints {
		if matches := predicate(res.GetMetadata()); !matches {
			Endpoints = append(Endpoints, res)
		}
	}
	s.Endpoints = Endpoints
	var Proxies gloo_solo_io.ProxyList
	for _, res := range s.Proxies {
		if matches := predicate(res.GetMetadata()); !matches {
			Proxies = append(Proxies, res)
		}
	}
	s.Proxies = Proxies
	var UpstreamGroups gloo_solo_io.UpstreamGroupList
	for _, res := range s.UpstreamGroups {
		if matches := predicate(res.GetMetadata()); !matches {
			UpstreamGroups = append(UpstreamGroups, res)
		}
	}
	s.UpstreamGroups = UpstreamGroups
	var Secrets gloo_solo_io.SecretList
	for _, res := range s.Secrets {
		if matches := predicate(res.GetMetadata()); !matches {
			Secrets = append(Secrets, res)
		}
	}
	s.Secrets = Secrets
	var Upstreams gloo_solo_io.UpstreamList
	for _, res := range s.Upstreams {
		if matches := predicate(res.GetMetadata()); !matches {
			Upstreams = append(Upstreams, res)
		}
	}
	s.Upstreams = Upstreams
	var AuthConfigs enterprise_gloo_solo_io.AuthConfigList
	for _, res := range s.AuthConfigs {
		if matches := predicate(res.GetMetadata()); !matches {
			AuthConfigs = append(AuthConfigs, res)
		}
	}
	s.AuthConfigs = AuthConfigs
	var Ratelimitconfigs github_com_solo_io_gloo_projects_controller_pkg_api_external_solo_ratelimit.RateLimitConfigList
	for _, res := range s.Ratelimitconfigs {
		if matches := predicate(res.GetMetadata()); !matches {
			Ratelimitconfigs = append(Ratelimitconfigs, res)
		}
	}
	s.Ratelimitconfigs = Ratelimitconfigs
	var VirtualServices gateway_solo_io.VirtualServiceList
	for _, res := range s.VirtualServices {
		if matches := predicate(res.GetMetadata()); !matches {
			VirtualServices = append(VirtualServices, res)
		}
	}
	s.VirtualServices = VirtualServices
	var RouteTables gateway_solo_io.RouteTableList
	for _, res := range s.RouteTables {
		if matches := predicate(res.GetMetadata()); !matches {
			RouteTables = append(RouteTables, res)
		}
	}
	s.RouteTables = RouteTables
	var Gateways gateway_solo_io.GatewayList
	for _, res := range s.Gateways {
		if matches := predicate(res.GetMetadata()); !matches {
			Gateways = append(Gateways, res)
		}
	}
	s.Gateways = Gateways
	var VirtualHostOptions gateway_solo_io.VirtualHostOptionList
	for _, res := range s.VirtualHostOptions {
		if matches := predicate(res.GetMetadata()); !matches {
			VirtualHostOptions = append(VirtualHostOptions, res)
		}
	}
	s.VirtualHostOptions = VirtualHostOptions
	var RouteOptions gateway_solo_io.RouteOptionList
	for _, res := range s.RouteOptions {
		if matches := predicate(res.GetMetadata()); !matches {
			RouteOptions = append(RouteOptions, res)
		}
	}
	s.RouteOptions = RouteOptions
	var HttpGateways gateway_solo_io.MatchableHttpGatewayList
	for _, res := range s.HttpGateways {
		if matches := predicate(res.GetMetadata()); !matches {
			HttpGateways = append(HttpGateways, res)
		}
	}
	s.HttpGateways = HttpGateways
	var TcpGateways gateway_solo_io.MatchableTcpGatewayList
	for _, res := range s.TcpGateways {
		if matches := predicate(res.GetMetadata()); !matches {
			TcpGateways = append(TcpGateways, res)
		}
	}
	s.TcpGateways = TcpGateways
	var GraphqlApis graphql_gloo_solo_io.GraphQLApiList
	for _, res := range s.GraphqlApis {
		if matches := predicate(res.GetMetadata()); !matches {
			GraphqlApis = append(GraphqlApis, res)
		}
	}
	s.GraphqlApis = GraphqlApis
}

func (s *ApiSnapshot) UpsertToResourceList(resource resources.Resource) error {
	refKey := resource.GetMetadata().Ref().Key()
	switch typed := resource.(type) {
	case *gloo_solo_io.Artifact:
		updated := false
		for i, res := range s.Artifacts {
			if refKey == res.GetMetadata().Ref().Key() {
				s.Artifacts[i] = typed
				updated = true
			}
		}
		if !updated {
			s.Artifacts = append(s.Artifacts, typed)
		}
		s.Artifacts.Sort()
		return nil
	case *gloo_solo_io.Endpoint:
		updated := false
		for i, res := range s.Endpoints {
			if refKey == res.GetMetadata().Ref().Key() {
				s.Endpoints[i] = typed
				updated = true
			}
		}
		if !updated {
			s.Endpoints = append(s.Endpoints, typed)
		}
		s.Endpoints.Sort()
		return nil
	case *gloo_solo_io.Proxy:
		updated := false
		for i, res := range s.Proxies {
			if refKey == res.GetMetadata().Ref().Key() {
				s.Proxies[i] = typed
				updated = true
			}
		}
		if !updated {
			s.Proxies = append(s.Proxies, typed)
		}
		s.Proxies.Sort()
		return nil
	case *gloo_solo_io.UpstreamGroup:
		updated := false
		for i, res := range s.UpstreamGroups {
			if refKey == res.GetMetadata().Ref().Key() {
				s.UpstreamGroups[i] = typed
				updated = true
			}
		}
		if !updated {
			s.UpstreamGroups = append(s.UpstreamGroups, typed)
		}
		s.UpstreamGroups.Sort()
		return nil
	case *gloo_solo_io.Secret:
		updated := false
		for i, res := range s.Secrets {
			if refKey == res.GetMetadata().Ref().Key() {
				s.Secrets[i] = typed
				updated = true
			}
		}
		if !updated {
			s.Secrets = append(s.Secrets, typed)
		}
		s.Secrets.Sort()
		return nil
	case *gloo_solo_io.Upstream:
		updated := false
		for i, res := range s.Upstreams {
			if refKey == res.GetMetadata().Ref().Key() {
				s.Upstreams[i] = typed
				updated = true
			}
		}
		if !updated {
			s.Upstreams = append(s.Upstreams, typed)
		}
		s.Upstreams.Sort()
		return nil
	case *enterprise_gloo_solo_io.AuthConfig:
		updated := false
		for i, res := range s.AuthConfigs {
			if refKey == res.GetMetadata().Ref().Key() {
				s.AuthConfigs[i] = typed
				updated = true
			}
		}
		if !updated {
			s.AuthConfigs = append(s.AuthConfigs, typed)
		}
		s.AuthConfigs.Sort()
		return nil
	case *github_com_solo_io_gloo_projects_controller_pkg_api_external_solo_ratelimit.RateLimitConfig:
		updated := false
		for i, res := range s.Ratelimitconfigs {
			if refKey == res.GetMetadata().Ref().Key() {
				s.Ratelimitconfigs[i] = typed
				updated = true
			}
		}
		if !updated {
			s.Ratelimitconfigs = append(s.Ratelimitconfigs, typed)
		}
		s.Ratelimitconfigs.Sort()
		return nil
	case *gateway_solo_io.VirtualService:
		updated := false
		for i, res := range s.VirtualServices {
			if refKey == res.GetMetadata().Ref().Key() {
				s.VirtualServices[i] = typed
				updated = true
			}
		}
		if !updated {
			s.VirtualServices = append(s.VirtualServices, typed)
		}
		s.VirtualServices.Sort()
		return nil
	case *gateway_solo_io.RouteTable:
		updated := false
		for i, res := range s.RouteTables {
			if refKey == res.GetMetadata().Ref().Key() {
				s.RouteTables[i] = typed
				updated = true
			}
		}
		if !updated {
			s.RouteTables = append(s.RouteTables, typed)
		}
		s.RouteTables.Sort()
		return nil
	case *gateway_solo_io.Gateway:
		updated := false
		for i, res := range s.Gateways {
			if refKey == res.GetMetadata().Ref().Key() {
				s.Gateways[i] = typed
				updated = true
			}
		}
		if !updated {
			s.Gateways = append(s.Gateways, typed)
		}
		s.Gateways.Sort()
		return nil
	case *gateway_solo_io.VirtualHostOption:
		updated := false
		for i, res := range s.VirtualHostOptions {
			if refKey == res.GetMetadata().Ref().Key() {
				s.VirtualHostOptions[i] = typed
				updated = true
			}
		}
		if !updated {
			s.VirtualHostOptions = append(s.VirtualHostOptions, typed)
		}
		s.VirtualHostOptions.Sort()
		return nil
	case *gateway_solo_io.RouteOption:
		updated := false
		for i, res := range s.RouteOptions {
			if refKey == res.GetMetadata().Ref().Key() {
				s.RouteOptions[i] = typed
				updated = true
			}
		}
		if !updated {
			s.RouteOptions = append(s.RouteOptions, typed)
		}
		s.RouteOptions.Sort()
		return nil
	case *gateway_solo_io.MatchableHttpGateway:
		updated := false
		for i, res := range s.HttpGateways {
			if refKey == res.GetMetadata().Ref().Key() {
				s.HttpGateways[i] = typed
				updated = true
			}
		}
		if !updated {
			s.HttpGateways = append(s.HttpGateways, typed)
		}
		s.HttpGateways.Sort()
		return nil
	case *gateway_solo_io.MatchableTcpGateway:
		updated := false
		for i, res := range s.TcpGateways {
			if refKey == res.GetMetadata().Ref().Key() {
				s.TcpGateways[i] = typed
				updated = true
			}
		}
		if !updated {
			s.TcpGateways = append(s.TcpGateways, typed)
		}
		s.TcpGateways.Sort()
		return nil
	case *graphql_gloo_solo_io.GraphQLApi:
		updated := false
		for i, res := range s.GraphqlApis {
			if refKey == res.GetMetadata().Ref().Key() {
				s.GraphqlApis[i] = typed
				updated = true
			}
		}
		if !updated {
			s.GraphqlApis = append(s.GraphqlApis, typed)
		}
		s.GraphqlApis.Sort()
		return nil
	default:
		return eris.Errorf("did not add/replace the resource type because it does not exist %T", resource)
	}
}

type ApiSnapshotStringer struct {
	Version            uint64
	Artifacts          []string
	Endpoints          []string
	Proxies            []string
	UpstreamGroups     []string
	Secrets            []string
	Upstreams          []string
	AuthConfigs        []string
	Ratelimitconfigs   []string
	VirtualServices    []string
	RouteTables        []string
	Gateways           []string
	VirtualHostOptions []string
	RouteOptions       []string
	HttpGateways       []string
	TcpGateways        []string
	GraphqlApis        []string
}

func (ss ApiSnapshotStringer) String() string {
	s := fmt.Sprintf("ApiSnapshot %v\n", ss.Version)

	s += fmt.Sprintf("  Artifacts %v\n", len(ss.Artifacts))
	for _, name := range ss.Artifacts {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Endpoints %v\n", len(ss.Endpoints))
	for _, name := range ss.Endpoints {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Proxies %v\n", len(ss.Proxies))
	for _, name := range ss.Proxies {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  UpstreamGroups %v\n", len(ss.UpstreamGroups))
	for _, name := range ss.UpstreamGroups {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Secrets %v\n", len(ss.Secrets))
	for _, name := range ss.Secrets {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Upstreams %v\n", len(ss.Upstreams))
	for _, name := range ss.Upstreams {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  AuthConfigs %v\n", len(ss.AuthConfigs))
	for _, name := range ss.AuthConfigs {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Ratelimitconfigs %v\n", len(ss.Ratelimitconfigs))
	for _, name := range ss.Ratelimitconfigs {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  VirtualServices %v\n", len(ss.VirtualServices))
	for _, name := range ss.VirtualServices {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  RouteTables %v\n", len(ss.RouteTables))
	for _, name := range ss.RouteTables {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Gateways %v\n", len(ss.Gateways))
	for _, name := range ss.Gateways {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  VirtualHostOptions %v\n", len(ss.VirtualHostOptions))
	for _, name := range ss.VirtualHostOptions {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  RouteOptions %v\n", len(ss.RouteOptions))
	for _, name := range ss.RouteOptions {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  HttpGateways %v\n", len(ss.HttpGateways))
	for _, name := range ss.HttpGateways {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  TcpGateways %v\n", len(ss.TcpGateways))
	for _, name := range ss.TcpGateways {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  GraphqlApis %v\n", len(ss.GraphqlApis))
	for _, name := range ss.GraphqlApis {
		s += fmt.Sprintf("    %v\n", name)
	}

	return s
}

func (s ApiSnapshot) Stringer() ApiSnapshotStringer {
	snapshotHash, err := s.Hash(nil)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	return ApiSnapshotStringer{
		Version:            snapshotHash,
		Artifacts:          s.Artifacts.NamespacesDotNames(),
		Endpoints:          s.Endpoints.NamespacesDotNames(),
		Proxies:            s.Proxies.NamespacesDotNames(),
		UpstreamGroups:     s.UpstreamGroups.NamespacesDotNames(),
		Secrets:            s.Secrets.NamespacesDotNames(),
		Upstreams:          s.Upstreams.NamespacesDotNames(),
		AuthConfigs:        s.AuthConfigs.NamespacesDotNames(),
		Ratelimitconfigs:   s.Ratelimitconfigs.NamespacesDotNames(),
		VirtualServices:    s.VirtualServices.NamespacesDotNames(),
		RouteTables:        s.RouteTables.NamespacesDotNames(),
		Gateways:           s.Gateways.NamespacesDotNames(),
		VirtualHostOptions: s.VirtualHostOptions.NamespacesDotNames(),
		RouteOptions:       s.RouteOptions.NamespacesDotNames(),
		HttpGateways:       s.HttpGateways.NamespacesDotNames(),
		TcpGateways:        s.TcpGateways.NamespacesDotNames(),
		GraphqlApis:        s.GraphqlApis.NamespacesDotNames(),
	}
}

var ApiGvkToHashableResource = map[schema.GroupVersionKind]func() resources.HashableResource{
	gloo_solo_io.ArtifactGVK:              gloo_solo_io.NewArtifactHashableResource,
	gloo_solo_io.EndpointGVK:              gloo_solo_io.NewEndpointHashableResource,
	gloo_solo_io.ProxyGVK:                 gloo_solo_io.NewProxyHashableResource,
	gloo_solo_io.UpstreamGroupGVK:         gloo_solo_io.NewUpstreamGroupHashableResource,
	gloo_solo_io.SecretGVK:                gloo_solo_io.NewSecretHashableResource,
	gloo_solo_io.UpstreamGVK:              gloo_solo_io.NewUpstreamHashableResource,
	enterprise_gloo_solo_io.AuthConfigGVK: enterprise_gloo_solo_io.NewAuthConfigHashableResource,
	github_com_solo_io_gloo_projects_controller_pkg_api_external_solo_ratelimit.RateLimitConfigGVK: github_com_solo_io_gloo_projects_controller_pkg_api_external_solo_ratelimit.NewRateLimitConfigHashableResource,
	gateway_solo_io.VirtualServiceGVK:       gateway_solo_io.NewVirtualServiceHashableResource,
	gateway_solo_io.RouteTableGVK:           gateway_solo_io.NewRouteTableHashableResource,
	gateway_solo_io.GatewayGVK:              gateway_solo_io.NewGatewayHashableResource,
	gateway_solo_io.VirtualHostOptionGVK:    gateway_solo_io.NewVirtualHostOptionHashableResource,
	gateway_solo_io.RouteOptionGVK:          gateway_solo_io.NewRouteOptionHashableResource,
	gateway_solo_io.MatchableHttpGatewayGVK: gateway_solo_io.NewMatchableHttpGatewayHashableResource,
	gateway_solo_io.MatchableTcpGatewayGVK:  gateway_solo_io.NewMatchableTcpGatewayHashableResource,
	graphql_gloo_solo_io.GraphQLApiGVK:      graphql_gloo_solo_io.NewGraphQLApiHashableResource,
}
