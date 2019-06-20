// Code generated from ./ASN.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // ASN

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseASNVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseASNVisitor) VisitModuleDefinition(ctx *ModuleDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitTagDefault(ctx *TagDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitExtensionDefault(ctx *ExtensionDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitModuleBody(ctx *ModuleBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitExports(ctx *ExportsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitSymbolsExported(ctx *SymbolsExportedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitImports(ctx *ImportsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitSymbolsImported(ctx *SymbolsImportedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitSymbolsFromModuleList(ctx *SymbolsFromModuleListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitSymbolsFromModule(ctx *SymbolsFromModuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitGlobalModuleReference(ctx *GlobalModuleReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitAssignedIdentifier(ctx *AssignedIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitSymbolList(ctx *SymbolListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitSymbol(ctx *SymbolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitAssignmentList(ctx *AssignmentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitSequenceType(ctx *SequenceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitExtensionAndException(ctx *ExtensionAndExceptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitOptionalExtensionMarker(ctx *OptionalExtensionMarkerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitComponentTypeLists(ctx *ComponentTypeListsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitRootComponentTypeList(ctx *RootComponentTypeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitComponentTypeList(ctx *ComponentTypeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitComponentType(ctx *ComponentTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitExtensionAdditions(ctx *ExtensionAdditionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitExtensionAdditionList(ctx *ExtensionAdditionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitExtensionAddition(ctx *ExtensionAdditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitExtensionAdditionGroup(ctx *ExtensionAdditionGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitVersionNumber(ctx *VersionNumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitSequenceOfType(ctx *SequenceOfTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitSizeConstraint(ctx *SizeConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitParameterizedAssignment(ctx *ParameterizedAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitParameterList(ctx *ParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitParameter(ctx *ParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitParamGovernor(ctx *ParamGovernorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitGovernor(ctx *GovernorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitObjectClassAssignment(ctx *ObjectClassAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitObjectClass(ctx *ObjectClassContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitDefinedObjectClass(ctx *DefinedObjectClassContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitUsefulObjectClassReference(ctx *UsefulObjectClassReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitExternalObjectClassReference(ctx *ExternalObjectClassReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitObjectClassDefn(ctx *ObjectClassDefnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitWithSyntaxSpec(ctx *WithSyntaxSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitSyntaxList(ctx *SyntaxListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitTokenOrGroupSpec(ctx *TokenOrGroupSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitOptionalGroup(ctx *OptionalGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitRequiredToken(ctx *RequiredTokenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitPrimitiveFieldName(ctx *PrimitiveFieldNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitFieldSpec(ctx *FieldSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitTypeFieldSpec(ctx *TypeFieldSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitTypeOptionalitySpec(ctx *TypeOptionalitySpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitFixedTypeValueFieldSpec(ctx *FixedTypeValueFieldSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitValueOptionalitySpec(ctx *ValueOptionalitySpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitVariableTypeValueFieldSpec(ctx *VariableTypeValueFieldSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitFixedTypeValueSetFieldSpec(ctx *FixedTypeValueSetFieldSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitValueSetOptionalitySpec(ctx *ValueSetOptionalitySpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitObject(ctx *ObjectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitParameterizedObject(ctx *ParameterizedObjectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitDefinedObject(ctx *DefinedObjectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitObjectSet(ctx *ObjectSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitObjectSetSpec(ctx *ObjectSetSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitFieldName(ctx *FieldNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitValueSet(ctx *ValueSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitElementSetSpecs(ctx *ElementSetSpecsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitRootElementSetSpec(ctx *RootElementSetSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitAdditionalElementSetSpec(ctx *AdditionalElementSetSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitElementSetSpec(ctx *ElementSetSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitUnions(ctx *UnionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitExclusions(ctx *ExclusionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitIntersections(ctx *IntersectionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitUnionMark(ctx *UnionMarkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitIntersectionMark(ctx *IntersectionMarkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitElements(ctx *ElementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitObjectSetElements(ctx *ObjectSetElementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitIntersectionElements(ctx *IntersectionElementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitSubtypeElements(ctx *SubtypeElementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitVariableTypeValueSetFieldSpec(ctx *VariableTypeValueSetFieldSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitObjectFieldSpec(ctx *ObjectFieldSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitObjectOptionalitySpec(ctx *ObjectOptionalitySpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitObjectSetFieldSpec(ctx *ObjectSetFieldSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitObjectSetOptionalitySpec(ctx *ObjectSetOptionalitySpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitTypeAssignment(ctx *TypeAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitValueAssignment(ctx *ValueAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitAsnType(ctx *AsnTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitBuiltinType(ctx *BuiltinTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitObjectClassFieldType(ctx *ObjectClassFieldTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitSetType(ctx *SetTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitSetOfType(ctx *SetOfTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitReferencedType(ctx *ReferencedTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitDefinedType(ctx *DefinedTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitConstraint(ctx *ConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitConstraintSpec(ctx *ConstraintSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitUserDefinedConstraint(ctx *UserDefinedConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitGeneralConstraint(ctx *GeneralConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitUserDefinedConstraintParameter(ctx *UserDefinedConstraintParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitTableConstraint(ctx *TableConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitSimpleTableConstraint(ctx *SimpleTableConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitContentsConstraint(ctx *ContentsConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitSubtypeConstraint(ctx *SubtypeConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitBuiltinValue(ctx *BuiltinValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitObjectIdentifierValue(ctx *ObjectIdentifierValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitObjIdComponentsList(ctx *ObjIdComponentsListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitObjIdComponents(ctx *ObjIdComponentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitIntegerValue(ctx *IntegerValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitChoiceValue(ctx *ChoiceValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitEnumeratedValue(ctx *EnumeratedValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitSignedNumber(ctx *SignedNumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitChoiceType(ctx *ChoiceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitAlternativeTypeLists(ctx *AlternativeTypeListsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitExtensionAdditionAlternatives(ctx *ExtensionAdditionAlternativesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitExtensionAdditionAlternativesList(ctx *ExtensionAdditionAlternativesListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitExtensionAdditionAlternative(ctx *ExtensionAdditionAlternativeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitExtensionAdditionAlternativesGroup(ctx *ExtensionAdditionAlternativesGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitRootAlternativeTypeList(ctx *RootAlternativeTypeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitAlternativeTypeList(ctx *AlternativeTypeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitNamedType(ctx *NamedTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitEnumeratedType(ctx *EnumeratedTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitEnumerations(ctx *EnumerationsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitRootEnumeration(ctx *RootEnumerationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitEnumeration(ctx *EnumerationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitEnumerationItem(ctx *EnumerationItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitNamedNumber(ctx *NamedNumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitDefinedValue(ctx *DefinedValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitParameterizedValue(ctx *ParameterizedValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitSimpleDefinedValue(ctx *SimpleDefinedValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitActualParameterList(ctx *ActualParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitActualParameter(ctx *ActualParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitExceptionSpec(ctx *ExceptionSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitExceptionIdentification(ctx *ExceptionIdentificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitAdditionalEnumeration(ctx *AdditionalEnumerationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitIntegerType(ctx *IntegerTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitNamedNumberList(ctx *NamedNumberListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitObjectidentifiertype(ctx *ObjectidentifiertypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitComponentRelationConstraint(ctx *ComponentRelationConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitAtNotation(ctx *AtNotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitLevel(ctx *LevelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitComponentIdList(ctx *ComponentIdListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitOctetStringType(ctx *OctetStringTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitBitStringType(ctx *BitStringTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitNamedBitList(ctx *NamedBitListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitNamedBit(ctx *NamedBitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseASNVisitor) VisitBooleanValue(ctx *BooleanValueContext) interface{} {
	return v.VisitChildren(ctx)
}
