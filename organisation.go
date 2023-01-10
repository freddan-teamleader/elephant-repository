package docformat

import (
	"context"
	"fmt"
)

func sideloadOrganisationInformation(
	ctx context.Context, oc ObjectGetter,
	evt ContentLogEvent, doc *Document,
) error {
	var con NMLConcept

	_, err := oc.GetObject(
		ctx, evt.UUID, evt.Content.Version, &con)
	if err != nil {
		return fmt.Errorf("failed to fetch original document: %w", err)
	}

	properties := make(map[string]string)

	for _, p := range con.Concept.Properties {
		properties[p.Type] = p.Literal
	}

	for _, p := range con.ItemMeta.Properties {
		properties[p.Type] = p.Literal
	}

	if properties["ttext:typ"] != "" {
		doc.Meta = withBlockOfType("tt/type", doc.Meta,
			func(block Block) Block {
				block.Value = properties["ttext:typ"]

				return block
			})
	}

	return nil
}
