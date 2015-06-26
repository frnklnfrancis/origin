package nodes

import (
	"fmt"
	"reflect"

	osgraph "github.com/openshift/origin/pkg/api/graph"
	imageapi "github.com/openshift/origin/pkg/image/api"
)

var (
	ImageStreamNodeKind    = reflect.TypeOf(imageapi.ImageStream{}).Name()
	ImageNodeKind          = reflect.TypeOf(imageapi.Image{}).Name()
	ImageStreamTagNodeKind = reflect.TypeOf(imageapi.ImageStreamTag{}).Name()

	// non-api types
	DockerRepositoryNodeKind = reflect.TypeOf(imageapi.DockerImageReference{}).Name()
	ImageLayerNodeKind       = "ImageLayer"
)

func ImageStreamNodeName(o *imageapi.ImageStream) osgraph.UniqueName {
	return osgraph.GetUniqueRuntimeObjectNodeName(ImageStreamNodeKind, o)
}

type ImageStreamNode struct {
	osgraph.Node
	*imageapi.ImageStream
}

func (n ImageStreamNode) Object() interface{} {
	return n.ImageStream
}

func (n ImageStreamNode) String() string {
	return string(ImageStreamNodeName(n.ImageStream))
}

func (*ImageStreamNode) Kind() string {
	return ImageStreamNodeKind
}

func ImageStreamTagNodeName(o *imageapi.ImageStreamTag) osgraph.UniqueName {
	return osgraph.GetUniqueRuntimeObjectNodeName(ImageStreamTagNodeKind, o)
}

type ImageStreamTagNode struct {
	osgraph.Node
	*imageapi.ImageStreamTag

	Synthetic bool
}

func (n ImageStreamTagNode) IsSynthetic() bool {
	return n.Synthetic
}

func (n ImageStreamTagNode) ImageSpec() string {
	name, tag, _ := imageapi.SplitImageStreamTag(n.ImageStreamTag.Name)
	return imageapi.DockerImageReference{Namespace: n.Namespace, Name: name, Tag: tag}.String()
}

func (n ImageStreamTagNode) ImageTag() string {
	_, tag, _ := imageapi.SplitImageStreamTag(n.ImageStreamTag.Name)
	return tag
}

func (n ImageStreamTagNode) Object() interface{} {
	return n.ImageStreamTag
}

func (n ImageStreamTagNode) String() string {
	return string(ImageStreamTagNodeName(n.ImageStreamTag))
}

func (*ImageStreamTagNode) Kind() string {
	return ImageStreamTagNodeKind
}

func DockerImageRepositoryNodeName(o imageapi.DockerImageReference) osgraph.UniqueName {
	return osgraph.UniqueName(fmt.Sprintf("%s|%s", DockerRepositoryNodeKind, o.String()))
}

type DockerImageRepositoryNode struct {
	osgraph.Node
	Ref imageapi.DockerImageReference
}

func (n DockerImageRepositoryNode) ImageSpec() string {
	return n.Ref.String()
}

func (n DockerImageRepositoryNode) ImageTag() string {
	return n.Ref.DockerClientDefaults().Tag
}

func (n DockerImageRepositoryNode) String() string {
	return string(DockerImageRepositoryNodeName(n.Ref))
}

func (*DockerImageRepositoryNode) Kind() string {
	return DockerRepositoryNodeKind
}

func ImageNodeName(o *imageapi.Image) osgraph.UniqueName {
	return osgraph.GetUniqueRuntimeObjectNodeName(ImageNodeKind, o)
}

type ImageNode struct {
	osgraph.Node
	Image *imageapi.Image
}

func (n ImageNode) Object() interface{} {
	return n.Image
}

func (n ImageNode) String() string {
	return string(ImageNodeName(n.Image))
}

func (*ImageNode) Kind() string {
	return ImageNodeKind
}

func ImageLayerNodeName(layer string) osgraph.UniqueName {
	return osgraph.UniqueName(fmt.Sprintf("%s|%s", ImageLayerNodeKind, layer))
}

type ImageLayerNode struct {
	osgraph.Node
	Layer string
}

func (n ImageLayerNode) Object() interface{} {
	return n.Layer
}

func (n ImageLayerNode) String() string {
	return string(ImageLayerNodeName(n.Layer))
}

func (*ImageLayerNode) Kind() string {
	return ImageLayerNodeKind
}
