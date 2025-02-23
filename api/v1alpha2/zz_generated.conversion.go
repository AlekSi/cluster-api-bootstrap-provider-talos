// +build !ignore_autogenerated

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha2

import (
	v1alpha3 "github.com/talos-systems/cluster-api-bootstrap-provider-talos/api/v1alpha3"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*TalosConfig)(nil), (*v1alpha3.TalosConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_TalosConfig_To_v1alpha3_TalosConfig(a.(*TalosConfig), b.(*v1alpha3.TalosConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha3.TalosConfig)(nil), (*TalosConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_TalosConfig_To_v1alpha2_TalosConfig(a.(*v1alpha3.TalosConfig), b.(*TalosConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*TalosConfigList)(nil), (*v1alpha3.TalosConfigList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_TalosConfigList_To_v1alpha3_TalosConfigList(a.(*TalosConfigList), b.(*v1alpha3.TalosConfigList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha3.TalosConfigList)(nil), (*TalosConfigList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_TalosConfigList_To_v1alpha2_TalosConfigList(a.(*v1alpha3.TalosConfigList), b.(*TalosConfigList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*TalosConfigTemplate)(nil), (*v1alpha3.TalosConfigTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_TalosConfigTemplate_To_v1alpha3_TalosConfigTemplate(a.(*TalosConfigTemplate), b.(*v1alpha3.TalosConfigTemplate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha3.TalosConfigTemplate)(nil), (*TalosConfigTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_TalosConfigTemplate_To_v1alpha2_TalosConfigTemplate(a.(*v1alpha3.TalosConfigTemplate), b.(*TalosConfigTemplate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*TalosConfigTemplateList)(nil), (*v1alpha3.TalosConfigTemplateList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_TalosConfigTemplateList_To_v1alpha3_TalosConfigTemplateList(a.(*TalosConfigTemplateList), b.(*v1alpha3.TalosConfigTemplateList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha3.TalosConfigTemplateList)(nil), (*TalosConfigTemplateList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_TalosConfigTemplateList_To_v1alpha2_TalosConfigTemplateList(a.(*v1alpha3.TalosConfigTemplateList), b.(*TalosConfigTemplateList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*TalosConfigTemplateResource)(nil), (*v1alpha3.TalosConfigTemplateResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_TalosConfigTemplateResource_To_v1alpha3_TalosConfigTemplateResource(a.(*TalosConfigTemplateResource), b.(*v1alpha3.TalosConfigTemplateResource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha3.TalosConfigTemplateResource)(nil), (*TalosConfigTemplateResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_TalosConfigTemplateResource_To_v1alpha2_TalosConfigTemplateResource(a.(*v1alpha3.TalosConfigTemplateResource), b.(*TalosConfigTemplateResource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*TalosConfigTemplateSpec)(nil), (*v1alpha3.TalosConfigTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_TalosConfigTemplateSpec_To_v1alpha3_TalosConfigTemplateSpec(a.(*TalosConfigTemplateSpec), b.(*v1alpha3.TalosConfigTemplateSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha3.TalosConfigTemplateSpec)(nil), (*TalosConfigTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_TalosConfigTemplateSpec_To_v1alpha2_TalosConfigTemplateSpec(a.(*v1alpha3.TalosConfigTemplateSpec), b.(*TalosConfigTemplateSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*TalosConfigSpec)(nil), (*v1alpha3.TalosConfigSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_TalosConfigSpec_To_v1alpha3_TalosConfigSpec(a.(*TalosConfigSpec), b.(*v1alpha3.TalosConfigSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*TalosConfigStatus)(nil), (*v1alpha3.TalosConfigStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_TalosConfigStatus_To_v1alpha3_TalosConfigStatus(a.(*TalosConfigStatus), b.(*v1alpha3.TalosConfigStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1alpha3.TalosConfigSpec)(nil), (*TalosConfigSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_TalosConfigSpec_To_v1alpha2_TalosConfigSpec(a.(*v1alpha3.TalosConfigSpec), b.(*TalosConfigSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1alpha3.TalosConfigStatus)(nil), (*TalosConfigStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_TalosConfigStatus_To_v1alpha2_TalosConfigStatus(a.(*v1alpha3.TalosConfigStatus), b.(*TalosConfigStatus), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha2_TalosConfig_To_v1alpha3_TalosConfig(in *TalosConfig, out *v1alpha3.TalosConfig, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha2_TalosConfigSpec_To_v1alpha3_TalosConfigSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha2_TalosConfigStatus_To_v1alpha3_TalosConfigStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha2_TalosConfig_To_v1alpha3_TalosConfig is an autogenerated conversion function.
func Convert_v1alpha2_TalosConfig_To_v1alpha3_TalosConfig(in *TalosConfig, out *v1alpha3.TalosConfig, s conversion.Scope) error {
	return autoConvert_v1alpha2_TalosConfig_To_v1alpha3_TalosConfig(in, out, s)
}

func autoConvert_v1alpha3_TalosConfig_To_v1alpha2_TalosConfig(in *v1alpha3.TalosConfig, out *TalosConfig, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha3_TalosConfigSpec_To_v1alpha2_TalosConfigSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha3_TalosConfigStatus_To_v1alpha2_TalosConfigStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha3_TalosConfig_To_v1alpha2_TalosConfig is an autogenerated conversion function.
func Convert_v1alpha3_TalosConfig_To_v1alpha2_TalosConfig(in *v1alpha3.TalosConfig, out *TalosConfig, s conversion.Scope) error {
	return autoConvert_v1alpha3_TalosConfig_To_v1alpha2_TalosConfig(in, out, s)
}

func autoConvert_v1alpha2_TalosConfigList_To_v1alpha3_TalosConfigList(in *TalosConfigList, out *v1alpha3.TalosConfigList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1alpha3.TalosConfig, len(*in))
		for i := range *in {
			if err := Convert_v1alpha2_TalosConfig_To_v1alpha3_TalosConfig(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha2_TalosConfigList_To_v1alpha3_TalosConfigList is an autogenerated conversion function.
func Convert_v1alpha2_TalosConfigList_To_v1alpha3_TalosConfigList(in *TalosConfigList, out *v1alpha3.TalosConfigList, s conversion.Scope) error {
	return autoConvert_v1alpha2_TalosConfigList_To_v1alpha3_TalosConfigList(in, out, s)
}

func autoConvert_v1alpha3_TalosConfigList_To_v1alpha2_TalosConfigList(in *v1alpha3.TalosConfigList, out *TalosConfigList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TalosConfig, len(*in))
		for i := range *in {
			if err := Convert_v1alpha3_TalosConfig_To_v1alpha2_TalosConfig(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha3_TalosConfigList_To_v1alpha2_TalosConfigList is an autogenerated conversion function.
func Convert_v1alpha3_TalosConfigList_To_v1alpha2_TalosConfigList(in *v1alpha3.TalosConfigList, out *TalosConfigList, s conversion.Scope) error {
	return autoConvert_v1alpha3_TalosConfigList_To_v1alpha2_TalosConfigList(in, out, s)
}

func autoConvert_v1alpha2_TalosConfigSpec_To_v1alpha3_TalosConfigSpec(in *TalosConfigSpec, out *v1alpha3.TalosConfigSpec, s conversion.Scope) error {
	out.GenerateType = in.GenerateType
	out.Data = in.Data
	return nil
}

func autoConvert_v1alpha3_TalosConfigSpec_To_v1alpha2_TalosConfigSpec(in *v1alpha3.TalosConfigSpec, out *TalosConfigSpec, s conversion.Scope) error {
	// WARNING: in.TalosVersion requires manual conversion: does not exist in peer-type
	out.GenerateType = in.GenerateType
	out.Data = in.Data
	// WARNING: in.ConfigPatches requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1alpha2_TalosConfigStatus_To_v1alpha3_TalosConfigStatus(in *TalosConfigStatus, out *v1alpha3.TalosConfigStatus, s conversion.Scope) error {
	out.Ready = in.Ready
	// WARNING: in.BootstrapData requires manual conversion: does not exist in peer-type
	out.TalosConfig = in.TalosConfig
	// WARNING: in.ErrorReason requires manual conversion: does not exist in peer-type
	// WARNING: in.ErrorMessage requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1alpha3_TalosConfigStatus_To_v1alpha2_TalosConfigStatus(in *v1alpha3.TalosConfigStatus, out *TalosConfigStatus, s conversion.Scope) error {
	out.Ready = in.Ready
	// WARNING: in.DataSecretName requires manual conversion: does not exist in peer-type
	out.TalosConfig = in.TalosConfig
	// WARNING: in.FailureReason requires manual conversion: does not exist in peer-type
	// WARNING: in.FailureMessage requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1alpha2_TalosConfigTemplate_To_v1alpha3_TalosConfigTemplate(in *TalosConfigTemplate, out *v1alpha3.TalosConfigTemplate, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha2_TalosConfigTemplateSpec_To_v1alpha3_TalosConfigTemplateSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha2_TalosConfigTemplate_To_v1alpha3_TalosConfigTemplate is an autogenerated conversion function.
func Convert_v1alpha2_TalosConfigTemplate_To_v1alpha3_TalosConfigTemplate(in *TalosConfigTemplate, out *v1alpha3.TalosConfigTemplate, s conversion.Scope) error {
	return autoConvert_v1alpha2_TalosConfigTemplate_To_v1alpha3_TalosConfigTemplate(in, out, s)
}

func autoConvert_v1alpha3_TalosConfigTemplate_To_v1alpha2_TalosConfigTemplate(in *v1alpha3.TalosConfigTemplate, out *TalosConfigTemplate, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha3_TalosConfigTemplateSpec_To_v1alpha2_TalosConfigTemplateSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha3_TalosConfigTemplate_To_v1alpha2_TalosConfigTemplate is an autogenerated conversion function.
func Convert_v1alpha3_TalosConfigTemplate_To_v1alpha2_TalosConfigTemplate(in *v1alpha3.TalosConfigTemplate, out *TalosConfigTemplate, s conversion.Scope) error {
	return autoConvert_v1alpha3_TalosConfigTemplate_To_v1alpha2_TalosConfigTemplate(in, out, s)
}

func autoConvert_v1alpha2_TalosConfigTemplateList_To_v1alpha3_TalosConfigTemplateList(in *TalosConfigTemplateList, out *v1alpha3.TalosConfigTemplateList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1alpha3.TalosConfigTemplate, len(*in))
		for i := range *in {
			if err := Convert_v1alpha2_TalosConfigTemplate_To_v1alpha3_TalosConfigTemplate(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha2_TalosConfigTemplateList_To_v1alpha3_TalosConfigTemplateList is an autogenerated conversion function.
func Convert_v1alpha2_TalosConfigTemplateList_To_v1alpha3_TalosConfigTemplateList(in *TalosConfigTemplateList, out *v1alpha3.TalosConfigTemplateList, s conversion.Scope) error {
	return autoConvert_v1alpha2_TalosConfigTemplateList_To_v1alpha3_TalosConfigTemplateList(in, out, s)
}

func autoConvert_v1alpha3_TalosConfigTemplateList_To_v1alpha2_TalosConfigTemplateList(in *v1alpha3.TalosConfigTemplateList, out *TalosConfigTemplateList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TalosConfigTemplate, len(*in))
		for i := range *in {
			if err := Convert_v1alpha3_TalosConfigTemplate_To_v1alpha2_TalosConfigTemplate(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha3_TalosConfigTemplateList_To_v1alpha2_TalosConfigTemplateList is an autogenerated conversion function.
func Convert_v1alpha3_TalosConfigTemplateList_To_v1alpha2_TalosConfigTemplateList(in *v1alpha3.TalosConfigTemplateList, out *TalosConfigTemplateList, s conversion.Scope) error {
	return autoConvert_v1alpha3_TalosConfigTemplateList_To_v1alpha2_TalosConfigTemplateList(in, out, s)
}

func autoConvert_v1alpha2_TalosConfigTemplateResource_To_v1alpha3_TalosConfigTemplateResource(in *TalosConfigTemplateResource, out *v1alpha3.TalosConfigTemplateResource, s conversion.Scope) error {
	if err := Convert_v1alpha2_TalosConfigSpec_To_v1alpha3_TalosConfigSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha2_TalosConfigTemplateResource_To_v1alpha3_TalosConfigTemplateResource is an autogenerated conversion function.
func Convert_v1alpha2_TalosConfigTemplateResource_To_v1alpha3_TalosConfigTemplateResource(in *TalosConfigTemplateResource, out *v1alpha3.TalosConfigTemplateResource, s conversion.Scope) error {
	return autoConvert_v1alpha2_TalosConfigTemplateResource_To_v1alpha3_TalosConfigTemplateResource(in, out, s)
}

func autoConvert_v1alpha3_TalosConfigTemplateResource_To_v1alpha2_TalosConfigTemplateResource(in *v1alpha3.TalosConfigTemplateResource, out *TalosConfigTemplateResource, s conversion.Scope) error {
	if err := Convert_v1alpha3_TalosConfigSpec_To_v1alpha2_TalosConfigSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha3_TalosConfigTemplateResource_To_v1alpha2_TalosConfigTemplateResource is an autogenerated conversion function.
func Convert_v1alpha3_TalosConfigTemplateResource_To_v1alpha2_TalosConfigTemplateResource(in *v1alpha3.TalosConfigTemplateResource, out *TalosConfigTemplateResource, s conversion.Scope) error {
	return autoConvert_v1alpha3_TalosConfigTemplateResource_To_v1alpha2_TalosConfigTemplateResource(in, out, s)
}

func autoConvert_v1alpha2_TalosConfigTemplateSpec_To_v1alpha3_TalosConfigTemplateSpec(in *TalosConfigTemplateSpec, out *v1alpha3.TalosConfigTemplateSpec, s conversion.Scope) error {
	if err := Convert_v1alpha2_TalosConfigTemplateResource_To_v1alpha3_TalosConfigTemplateResource(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha2_TalosConfigTemplateSpec_To_v1alpha3_TalosConfigTemplateSpec is an autogenerated conversion function.
func Convert_v1alpha2_TalosConfigTemplateSpec_To_v1alpha3_TalosConfigTemplateSpec(in *TalosConfigTemplateSpec, out *v1alpha3.TalosConfigTemplateSpec, s conversion.Scope) error {
	return autoConvert_v1alpha2_TalosConfigTemplateSpec_To_v1alpha3_TalosConfigTemplateSpec(in, out, s)
}

func autoConvert_v1alpha3_TalosConfigTemplateSpec_To_v1alpha2_TalosConfigTemplateSpec(in *v1alpha3.TalosConfigTemplateSpec, out *TalosConfigTemplateSpec, s conversion.Scope) error {
	if err := Convert_v1alpha3_TalosConfigTemplateResource_To_v1alpha2_TalosConfigTemplateResource(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha3_TalosConfigTemplateSpec_To_v1alpha2_TalosConfigTemplateSpec is an autogenerated conversion function.
func Convert_v1alpha3_TalosConfigTemplateSpec_To_v1alpha2_TalosConfigTemplateSpec(in *v1alpha3.TalosConfigTemplateSpec, out *TalosConfigTemplateSpec, s conversion.Scope) error {
	return autoConvert_v1alpha3_TalosConfigTemplateSpec_To_v1alpha2_TalosConfigTemplateSpec(in, out, s)
}
