package ibm

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/IBM-Bluemix/bluemix-go/api/iampap/iampapv1"
	"github.com/IBM-Bluemix/bluemix-go/api/mccp/mccpv2"
	"github.com/apache/incubator-openwhisk-client-go/whisk"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/softlayer/softlayer-go/datatypes"
)

//HashInt ...
func HashInt(v interface{}) int { return v.(int) }

func expandStringList(input []interface{}) []string {
	vs := make([]string, len(input))
	for i, v := range input {
		vs[i] = v.(string)
	}
	return vs
}

func flattenStringList(list []string) []interface{} {
	vs := make([]interface{}, len(list))
	for i, v := range list {
		vs[i] = v
	}
	return vs
}

func expandIntList(input []interface{}) []int {
	vs := make([]int, len(input))
	for i, v := range input {
		vs[i] = v.(int)
	}
	return vs
}

func flattenIntList(list []int) []interface{} {
	vs := make([]interface{}, len(list))
	for i, v := range list {
		vs[i] = v
	}
	return vs
}

func newStringSet(f schema.SchemaSetFunc, in []string) *schema.Set {
	var out = make([]interface{}, len(in), len(in))
	for i, v := range in {
		out[i] = v
	}
	return schema.NewSet(f, out)
}

func flattenRoute(in []mccpv2.Route) *schema.Set {
	vs := make([]string, len(in))
	for i, v := range in {
		vs[i] = v.GUID
	}
	return newStringSet(schema.HashString, vs)
}

func stringSliceToSet(in []string) *schema.Set {
	vs := make([]string, len(in))
	for i, v := range in {
		vs[i] = v
	}
	return newStringSet(schema.HashString, vs)
}

func flattenServiceBindings(in []mccpv2.ServiceBinding) *schema.Set {
	vs := make([]string, len(in))
	for i, v := range in {
		vs[i] = v.ServiceInstanceGUID
	}
	return newStringSet(schema.HashString, vs)
}

func flattenPort(in []int) *schema.Set {
	var out = make([]interface{}, len(in))
	for i, v := range in {
		out[i] = v
	}
	return schema.NewSet(HashInt, out)
}

func flattenFileStorageID(in []datatypes.Network_Storage) *schema.Set {
	var out = []interface{}{}
	for _, v := range in {
		if *v.NasType == "NAS" {
			out = append(out, *v.Id)
		}
	}
	return schema.NewSet(HashInt, out)
}

func flattenBlockStorageID(in []datatypes.Network_Storage) *schema.Set {
	var out = []interface{}{}
	for _, v := range in {
		if *v.NasType == "ISCSI" {
			out = append(out, *v.Id)
		}
	}
	return schema.NewSet(HashInt, out)
}

func flattenSpaceRoleUsers(in []mccpv2.SpaceRole) *schema.Set {
	var out = []interface{}{}
	for _, v := range in {
		out = append(out, v.UserName)
	}
	return schema.NewSet(schema.HashString, out)
}

func flattenMapInterfaceVal(m map[string]interface{}) map[string]string {
	out := make(map[string]string)
	for k, v := range m {
		out[k] = fmt.Sprintf("%v", v)
	}
	return out
}

func flattenCredentials(creds map[string]interface{}) map[string]string {
	return flattenMapInterfaceVal(creds)
}

func flattenServiceKeyCredentials(creds map[string]interface{}) map[string]string {
	return flattenCredentials(creds)
}

func flattenServiceInstanceCredentials(keys []mccpv2.ServiceKeyFields) []interface{} {
	var out = make([]interface{}, len(keys), len(keys))
	for i, k := range keys {
		m := make(map[string]interface{})
		m["name"] = k.Entity.Name
		m["credentials"] = flattenServiceKeyCredentials(k.Entity.Credentials)
		out[i] = m
	}
	return out
}

func flattenIAMPolicyResource(list []iampapv1.Resources, iamClient iampapv1.IAMPAPAPI) ([]map[string]interface{}, error) {
	result := make([]map[string]interface{}, 0, len(list))
	for _, i := range list {
		name := i.ServiceName
		if name == "" {
			name = allIAMEnabledServices
		}
		serviceName, err := iamClient.IAMService().GetServiceDispalyName(name)
		if err != nil {
			return result, fmt.Errorf("Error retrieving service : %s", err)
		}
		l := map[string]interface{}{
			"service_name":      serviceName,
			"region":            i.Region,
			"resource_type":     i.ResourceType,
			"resource":          i.Resource,
			"space_guid":        i.SpaceId,
			"organization_guid": i.OrganizationId,
		}
		if i.ServiceInstance != "" {
			l["service_instance"] = []string{i.ServiceInstance}
		}
		result = append(result, l)
	}
	return result, nil
}

func flattenIAMPolicyRoles(list []iampapv1.Roles) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(list))
	for _, v := range list {
		l := map[string]interface{}{
			"name": roleIDToName[v.ID],
		}
		result = append(result, l)
	}
	return result
}

func normalizeJSONString(jsonString interface{}) (string, error) {
	var j interface{}
	if jsonString == nil || jsonString.(string) == "" {
		return "", nil
	}
	s := jsonString.(string)
	err := json.Unmarshal([]byte(s), &j)
	if err != nil {
		return s, err
	}
	bytes, err := json.Marshal(j)
	if err != nil {
		return "", err
	}
	return string(bytes[:]), nil
}

func expandLimits(l []interface{}) *whisk.Limits {
	if len(l) == 0 || l[0] == nil {
		return &whisk.Limits{}
	}
	in := l[0].(map[string]interface{})
	obj := &whisk.Limits{
		Timeout: ptrToInt(in["timeout"].(int)),
		Memory:  ptrToInt(in["memory"].(int)),
		Logsize: ptrToInt(in["log_size"].(int)),
	}
	return obj
}

func flattenLimits(in *whisk.Limits) []interface{} {
	att := make(map[string]interface{})
	if in.Timeout != nil {
		att["timeout"] = *in.Timeout
	}
	if in.Memory != nil {
		att["memory"] = *in.Memory
	}
	if in.Memory != nil {
		att["log_size"] = *in.Logsize
	}
	return []interface{}{att}
}

func expandExec(execs *schema.Set) *whisk.Exec {
	for _, exec := range execs.List() {
		e, _ := exec.(map[string]interface{})
		obj := &whisk.Exec{
			Image:      e["image"].(string),
			Init:       e["init"].(string),
			Code:       ptrToString(e["code"].(string)),
			Kind:       e["kind"].(string),
			Main:       e["main"].(string),
			Components: expandStringList(e["components"].([]interface{})),
		}
		return obj
	}

	return &whisk.Exec{}
}

func flattenExec(in *whisk.Exec) []interface{} {
	att := make(map[string]interface{})
	if in.Image != "" {
		att["image"] = in.Image
	}
	if in.Init != "" {
		att["init"] = in.Init
	}
	if in.Code != nil {
		att["code"] = *in.Code
	}
	if in.Kind != "" {
		att["kind"] = in.Kind
	}
	if in.Main != "" {
		att["main"] = in.Main
	}

	if len(in.Components) > 0 {
		att["components"] = flattenStringList(in.Components)
	}

	return []interface{}{att}
}

func expandAnnotations(annotations string) (whisk.KeyValueArr, error) {
	var result whisk.KeyValueArr
	dc := json.NewDecoder(strings.NewReader(annotations))
	dc.UseNumber()
	err := dc.Decode(&result)
	return result, err
}

func flattenAnnotations(in whisk.KeyValueArr) (string, error) {
	noExec := make(whisk.KeyValueArr, 0, len(in))
	for _, v := range in {
		if v.Key == "exec" {
			continue
		}
		noExec = append(noExec, v)
	}
	b, err := json.Marshal(noExec)
	if err != nil {
		return "", err
	}
	return string(b[:]), nil
}

func expandParameters(annotations string) (whisk.KeyValueArr, error) {
	var result whisk.KeyValueArr
	dc := json.NewDecoder(strings.NewReader(annotations))
	dc.UseNumber()
	err := dc.Decode(&result)
	return result, err
}

func flattenParameters(in whisk.KeyValueArr) (string, error) {
	noExec := make(whisk.KeyValueArr, 0, len(in))
	for _, v := range in {
		if v.Key == "_actions" {
			continue
		}
		noExec = append(noExec, v)
	}
	b, err := json.Marshal(noExec)
	if err != nil {
		return "", err
	}
	return string(b[:]), nil
}

func expandBinding(binding []interface{}) (*whisk.Binding, error) {
	bindingParameters := &whisk.Binding{}
	for _, b := range binding {
		param := b.(map[string]interface{})
		var qualifiedBindingName = new(QualifiedName)
		var err error
		bindingName := fmt.Sprintf("/%s/%s", param["namespace"].(string), param["name"].(string))
		if qualifiedBindingName, err = NewQualifiedName(bindingName); err != nil {
			return nil, NewQualifiedNameError(bindingName, err)
		}
		bindingParameters.Namespace = qualifiedBindingName.GetNamespace()
		bindingParameters.Name = qualifiedBindingName.GetEntityName()
	}

	return bindingParameters, nil
}

func flattenBinding(bindingParameters *whisk.Binding) []map[string]interface{} {
	binding := make(map[string]interface{})
	binding["namespace"] = bindingParameters.Namespace
	binding["name"] = bindingParameters.Name
	result := []map[string]interface{}{binding}
	return result
}

func flattenActionOrTrigger(in interface{}) []interface{} {
	att := make(map[string]interface{})
	parameter := in.(map[string]interface{})
	att["name"] = parameter["name"]
	att["path"] = parameter["path"]
	return []interface{}{att}
}

func flattenAction(namespace, name string) []interface{} {
	att := make(map[string]interface{})
	att["name"] = name
	packageName := ""
	temp := strings.Split(namespace, "/")
	if len(temp) == 2 {
		packageName = temp[1]
	}
	att["package"] = packageName
	return []interface{}{att}
}

func ptrToInt(i int) *int {
	return &i
}

func ptrToString(s string) *string {
	return &s
}

func filterAnnotations(bindedAnnotations, annotations whisk.KeyValueArr) whisk.KeyValueArr {
	userDefinedAnnotations := make(whisk.KeyValueArr, 0)
	for _, a := range annotations {
		insert := false
		if a.Key == "binding" {
			insert = false
			break
		}
		for _, b := range bindedAnnotations {
			if a.Key == b.Key && reflect.DeepEqual(a.Value, b.Value) {
				insert = false
				break
			}
			insert = true
		}
		if insert {
			userDefinedAnnotations = append(userDefinedAnnotations, a)
		}
	}
	return userDefinedAnnotations
}

func filterParameters(bindedParameters, parameters whisk.KeyValueArr) whisk.KeyValueArr {
	userDefinedParameters := make(whisk.KeyValueArr, 0)
	for _, p := range parameters {
		insert := false
		for _, b := range bindedParameters {
			if p.Key == b.Key && reflect.DeepEqual(p.Value, b.Value) {
				insert = false
				break
			}
			insert = true
		}
		if insert {
			userDefinedParameters = append(userDefinedParameters, p)
		}

	}
	return userDefinedParameters
}
