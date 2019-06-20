// Code generated from ./ASN.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // ASN

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by ASNParser.
type ASNVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ASNParser#moduleDefinition.
	VisitModuleDefinition(ctx *ModuleDefinitionContext) interface{}

	// Visit a parse tree produced by ASNParser#tagDefault.
	VisitTagDefault(ctx *TagDefaultContext) interface{}

	// Visit a parse tree produced by ASNParser#extensionDefault.
	VisitExtensionDefault(ctx *ExtensionDefaultContext) interface{}

	// Visit a parse tree produced by ASNParser#moduleBody.
	VisitModuleBody(ctx *ModuleBodyContext) interface{}

	// Visit a parse tree produced by ASNParser#exports.
	VisitExports(ctx *ExportsContext) interface{}

	// Visit a parse tree produced by ASNParser#symbolsExported.
	VisitSymbolsExported(ctx *SymbolsExportedContext) interface{}

	// Visit a parse tree produced by ASNParser#imports.
	VisitImports(ctx *ImportsContext) interface{}

	// Visit a parse tree produced by ASNParser#symbolsImported.
	VisitSymbolsImported(ctx *SymbolsImportedContext) interface{}

	// Visit a parse tree produced by ASNParser#symbolsFromModuleList.
	VisitSymbolsFromModuleList(ctx *SymbolsFromModuleListContext) interface{}

	// Visit a parse tree produced by ASNParser#symbolsFromModule.
	VisitSymbolsFromModule(ctx *SymbolsFromModuleContext) interface{}

	// Visit a parse tree produced by ASNParser#globalModuleReference.
	VisitGlobalModuleReference(ctx *GlobalModuleReferenceContext) interface{}

	// Visit a parse tree produced by ASNParser#assignedIdentifier.
	VisitAssignedIdentifier(ctx *AssignedIdentifierContext) interface{}

	// Visit a parse tree produced by ASNParser#symbolList.
	VisitSymbolList(ctx *SymbolListContext) interface{}

	// Visit a parse tree produced by ASNParser#symbol.
	VisitSymbol(ctx *SymbolContext) interface{}

	// Visit a parse tree produced by ASNParser#assignmentList.
	VisitAssignmentList(ctx *AssignmentListContext) interface{}

	// Visit a parse tree produced by ASNParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by ASNParser#sequenceType.
	VisitSequenceType(ctx *SequenceTypeContext) interface{}

	// Visit a parse tree produced by ASNParser#extensionAndException.
	VisitExtensionAndException(ctx *ExtensionAndExceptionContext) interface{}

	// Visit a parse tree produced by ASNParser#optionalExtensionMarker.
	VisitOptionalExtensionMarker(ctx *OptionalExtensionMarkerContext) interface{}

	// Visit a parse tree produced by ASNParser#componentTypeLists.
	VisitComponentTypeLists(ctx *ComponentTypeListsContext) interface{}

	// Visit a parse tree produced by ASNParser#rootComponentTypeList.
	VisitRootComponentTypeList(ctx *RootComponentTypeListContext) interface{}

	// Visit a parse tree produced by ASNParser#componentTypeList.
	VisitComponentTypeList(ctx *ComponentTypeListContext) interface{}

	// Visit a parse tree produced by ASNParser#componentType.
	VisitComponentType(ctx *ComponentTypeContext) interface{}

	// Visit a parse tree produced by ASNParser#extensionAdditions.
	VisitExtensionAdditions(ctx *ExtensionAdditionsContext) interface{}

	// Visit a parse tree produced by ASNParser#extensionAdditionList.
	VisitExtensionAdditionList(ctx *ExtensionAdditionListContext) interface{}

	// Visit a parse tree produced by ASNParser#extensionAddition.
	VisitExtensionAddition(ctx *ExtensionAdditionContext) interface{}

	// Visit a parse tree produced by ASNParser#extensionAdditionGroup.
	VisitExtensionAdditionGroup(ctx *ExtensionAdditionGroupContext) interface{}

	// Visit a parse tree produced by ASNParser#versionNumber.
	VisitVersionNumber(ctx *VersionNumberContext) interface{}

	// Visit a parse tree produced by ASNParser#sequenceOfType.
	VisitSequenceOfType(ctx *SequenceOfTypeContext) interface{}

	// Visit a parse tree produced by ASNParser#sizeConstraint.
	VisitSizeConstraint(ctx *SizeConstraintContext) interface{}

	// Visit a parse tree produced by ASNParser#parameterizedAssignment.
	VisitParameterizedAssignment(ctx *ParameterizedAssignmentContext) interface{}

	// Visit a parse tree produced by ASNParser#parameterList.
	VisitParameterList(ctx *ParameterListContext) interface{}

	// Visit a parse tree produced by ASNParser#parameter.
	VisitParameter(ctx *ParameterContext) interface{}

	// Visit a parse tree produced by ASNParser#paramGovernor.
	VisitParamGovernor(ctx *ParamGovernorContext) interface{}

	// Visit a parse tree produced by ASNParser#governor.
	VisitGovernor(ctx *GovernorContext) interface{}

	// Visit a parse tree produced by ASNParser#objectClassAssignment.
	VisitObjectClassAssignment(ctx *ObjectClassAssignmentContext) interface{}

	// Visit a parse tree produced by ASNParser#objectClass.
	VisitObjectClass(ctx *ObjectClassContext) interface{}

	// Visit a parse tree produced by ASNParser#definedObjectClass.
	VisitDefinedObjectClass(ctx *DefinedObjectClassContext) interface{}

	// Visit a parse tree produced by ASNParser#usefulObjectClassReference.
	VisitUsefulObjectClassReference(ctx *UsefulObjectClassReferenceContext) interface{}

	// Visit a parse tree produced by ASNParser#externalObjectClassReference.
	VisitExternalObjectClassReference(ctx *ExternalObjectClassReferenceContext) interface{}

	// Visit a parse tree produced by ASNParser#objectClassDefn.
	VisitObjectClassDefn(ctx *ObjectClassDefnContext) interface{}

	// Visit a parse tree produced by ASNParser#withSyntaxSpec.
	VisitWithSyntaxSpec(ctx *WithSyntaxSpecContext) interface{}

	// Visit a parse tree produced by ASNParser#syntaxList.
	VisitSyntaxList(ctx *SyntaxListContext) interface{}

	// Visit a parse tree produced by ASNParser#tokenOrGroupSpec.
	VisitTokenOrGroupSpec(ctx *TokenOrGroupSpecContext) interface{}

	// Visit a parse tree produced by ASNParser#optionalGroup.
	VisitOptionalGroup(ctx *OptionalGroupContext) interface{}

	// Visit a parse tree produced by ASNParser#requiredToken.
	VisitRequiredToken(ctx *RequiredTokenContext) interface{}

	// Visit a parse tree produced by ASNParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by ASNParser#primitiveFieldName.
	VisitPrimitiveFieldName(ctx *PrimitiveFieldNameContext) interface{}

	// Visit a parse tree produced by ASNParser#fieldSpec.
	VisitFieldSpec(ctx *FieldSpecContext) interface{}

	// Visit a parse tree produced by ASNParser#typeFieldSpec.
	VisitTypeFieldSpec(ctx *TypeFieldSpecContext) interface{}

	// Visit a parse tree produced by ASNParser#typeOptionalitySpec.
	VisitTypeOptionalitySpec(ctx *TypeOptionalitySpecContext) interface{}

	// Visit a parse tree produced by ASNParser#fixedTypeValueFieldSpec.
	VisitFixedTypeValueFieldSpec(ctx *FixedTypeValueFieldSpecContext) interface{}

	// Visit a parse tree produced by ASNParser#valueOptionalitySpec.
	VisitValueOptionalitySpec(ctx *ValueOptionalitySpecContext) interface{}

	// Visit a parse tree produced by ASNParser#variableTypeValueFieldSpec.
	VisitVariableTypeValueFieldSpec(ctx *VariableTypeValueFieldSpecContext) interface{}

	// Visit a parse tree produced by ASNParser#fixedTypeValueSetFieldSpec.
	VisitFixedTypeValueSetFieldSpec(ctx *FixedTypeValueSetFieldSpecContext) interface{}

	// Visit a parse tree produced by ASNParser#valueSetOptionalitySpec.
	VisitValueSetOptionalitySpec(ctx *ValueSetOptionalitySpecContext) interface{}

	// Visit a parse tree produced by ASNParser#object.
	VisitObject(ctx *ObjectContext) interface{}

	// Visit a parse tree produced by ASNParser#parameterizedObject.
	VisitParameterizedObject(ctx *ParameterizedObjectContext) interface{}

	// Visit a parse tree produced by ASNParser#definedObject.
	VisitDefinedObject(ctx *DefinedObjectContext) interface{}

	// Visit a parse tree produced by ASNParser#objectSet.
	VisitObjectSet(ctx *ObjectSetContext) interface{}

	// Visit a parse tree produced by ASNParser#objectSetSpec.
	VisitObjectSetSpec(ctx *ObjectSetSpecContext) interface{}

	// Visit a parse tree produced by ASNParser#fieldName.
	VisitFieldName(ctx *FieldNameContext) interface{}

	// Visit a parse tree produced by ASNParser#valueSet.
	VisitValueSet(ctx *ValueSetContext) interface{}

	// Visit a parse tree produced by ASNParser#elementSetSpecs.
	VisitElementSetSpecs(ctx *ElementSetSpecsContext) interface{}

	// Visit a parse tree produced by ASNParser#rootElementSetSpec.
	VisitRootElementSetSpec(ctx *RootElementSetSpecContext) interface{}

	// Visit a parse tree produced by ASNParser#additionalElementSetSpec.
	VisitAdditionalElementSetSpec(ctx *AdditionalElementSetSpecContext) interface{}

	// Visit a parse tree produced by ASNParser#elementSetSpec.
	VisitElementSetSpec(ctx *ElementSetSpecContext) interface{}

	// Visit a parse tree produced by ASNParser#unions.
	VisitUnions(ctx *UnionsContext) interface{}

	// Visit a parse tree produced by ASNParser#exclusions.
	VisitExclusions(ctx *ExclusionsContext) interface{}

	// Visit a parse tree produced by ASNParser#intersections.
	VisitIntersections(ctx *IntersectionsContext) interface{}

	// Visit a parse tree produced by ASNParser#unionMark.
	VisitUnionMark(ctx *UnionMarkContext) interface{}

	// Visit a parse tree produced by ASNParser#intersectionMark.
	VisitIntersectionMark(ctx *IntersectionMarkContext) interface{}

	// Visit a parse tree produced by ASNParser#elements.
	VisitElements(ctx *ElementsContext) interface{}

	// Visit a parse tree produced by ASNParser#objectSetElements.
	VisitObjectSetElements(ctx *ObjectSetElementsContext) interface{}

	// Visit a parse tree produced by ASNParser#intersectionElements.
	VisitIntersectionElements(ctx *IntersectionElementsContext) interface{}

	// Visit a parse tree produced by ASNParser#subtypeElements.
	VisitSubtypeElements(ctx *SubtypeElementsContext) interface{}

	// Visit a parse tree produced by ASNParser#variableTypeValueSetFieldSpec.
	VisitVariableTypeValueSetFieldSpec(ctx *VariableTypeValueSetFieldSpecContext) interface{}

	// Visit a parse tree produced by ASNParser#objectFieldSpec.
	VisitObjectFieldSpec(ctx *ObjectFieldSpecContext) interface{}

	// Visit a parse tree produced by ASNParser#objectOptionalitySpec.
	VisitObjectOptionalitySpec(ctx *ObjectOptionalitySpecContext) interface{}

	// Visit a parse tree produced by ASNParser#objectSetFieldSpec.
	VisitObjectSetFieldSpec(ctx *ObjectSetFieldSpecContext) interface{}

	// Visit a parse tree produced by ASNParser#objectSetOptionalitySpec.
	VisitObjectSetOptionalitySpec(ctx *ObjectSetOptionalitySpecContext) interface{}

	// Visit a parse tree produced by ASNParser#typeAssignment.
	VisitTypeAssignment(ctx *TypeAssignmentContext) interface{}

	// Visit a parse tree produced by ASNParser#valueAssignment.
	VisitValueAssignment(ctx *ValueAssignmentContext) interface{}

	// Visit a parse tree produced by ASNParser#asnType.
	VisitAsnType(ctx *AsnTypeContext) interface{}

	// Visit a parse tree produced by ASNParser#builtinType.
	VisitBuiltinType(ctx *BuiltinTypeContext) interface{}

	// Visit a parse tree produced by ASNParser#objectClassFieldType.
	VisitObjectClassFieldType(ctx *ObjectClassFieldTypeContext) interface{}

	// Visit a parse tree produced by ASNParser#setType.
	VisitSetType(ctx *SetTypeContext) interface{}

	// Visit a parse tree produced by ASNParser#setOfType.
	VisitSetOfType(ctx *SetOfTypeContext) interface{}

	// Visit a parse tree produced by ASNParser#referencedType.
	VisitReferencedType(ctx *ReferencedTypeContext) interface{}

	// Visit a parse tree produced by ASNParser#definedType.
	VisitDefinedType(ctx *DefinedTypeContext) interface{}

	// Visit a parse tree produced by ASNParser#constraint.
	VisitConstraint(ctx *ConstraintContext) interface{}

	// Visit a parse tree produced by ASNParser#constraintSpec.
	VisitConstraintSpec(ctx *ConstraintSpecContext) interface{}

	// Visit a parse tree produced by ASNParser#userDefinedConstraint.
	VisitUserDefinedConstraint(ctx *UserDefinedConstraintContext) interface{}

	// Visit a parse tree produced by ASNParser#generalConstraint.
	VisitGeneralConstraint(ctx *GeneralConstraintContext) interface{}

	// Visit a parse tree produced by ASNParser#userDefinedConstraintParameter.
	VisitUserDefinedConstraintParameter(ctx *UserDefinedConstraintParameterContext) interface{}

	// Visit a parse tree produced by ASNParser#tableConstraint.
	VisitTableConstraint(ctx *TableConstraintContext) interface{}

	// Visit a parse tree produced by ASNParser#simpleTableConstraint.
	VisitSimpleTableConstraint(ctx *SimpleTableConstraintContext) interface{}

	// Visit a parse tree produced by ASNParser#contentsConstraint.
	VisitContentsConstraint(ctx *ContentsConstraintContext) interface{}

	// Visit a parse tree produced by ASNParser#subtypeConstraint.
	VisitSubtypeConstraint(ctx *SubtypeConstraintContext) interface{}

	// Visit a parse tree produced by ASNParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by ASNParser#builtinValue.
	VisitBuiltinValue(ctx *BuiltinValueContext) interface{}

	// Visit a parse tree produced by ASNParser#objectIdentifierValue.
	VisitObjectIdentifierValue(ctx *ObjectIdentifierValueContext) interface{}

	// Visit a parse tree produced by ASNParser#objIdComponentsList.
	VisitObjIdComponentsList(ctx *ObjIdComponentsListContext) interface{}

	// Visit a parse tree produced by ASNParser#objIdComponents.
	VisitObjIdComponents(ctx *ObjIdComponentsContext) interface{}

	// Visit a parse tree produced by ASNParser#integerValue.
	VisitIntegerValue(ctx *IntegerValueContext) interface{}

	// Visit a parse tree produced by ASNParser#choiceValue.
	VisitChoiceValue(ctx *ChoiceValueContext) interface{}

	// Visit a parse tree produced by ASNParser#enumeratedValue.
	VisitEnumeratedValue(ctx *EnumeratedValueContext) interface{}

	// Visit a parse tree produced by ASNParser#signedNumber.
	VisitSignedNumber(ctx *SignedNumberContext) interface{}

	// Visit a parse tree produced by ASNParser#choiceType.
	VisitChoiceType(ctx *ChoiceTypeContext) interface{}

	// Visit a parse tree produced by ASNParser#alternativeTypeLists.
	VisitAlternativeTypeLists(ctx *AlternativeTypeListsContext) interface{}

	// Visit a parse tree produced by ASNParser#extensionAdditionAlternatives.
	VisitExtensionAdditionAlternatives(ctx *ExtensionAdditionAlternativesContext) interface{}

	// Visit a parse tree produced by ASNParser#extensionAdditionAlternativesList.
	VisitExtensionAdditionAlternativesList(ctx *ExtensionAdditionAlternativesListContext) interface{}

	// Visit a parse tree produced by ASNParser#extensionAdditionAlternative.
	VisitExtensionAdditionAlternative(ctx *ExtensionAdditionAlternativeContext) interface{}

	// Visit a parse tree produced by ASNParser#extensionAdditionAlternativesGroup.
	VisitExtensionAdditionAlternativesGroup(ctx *ExtensionAdditionAlternativesGroupContext) interface{}

	// Visit a parse tree produced by ASNParser#rootAlternativeTypeList.
	VisitRootAlternativeTypeList(ctx *RootAlternativeTypeListContext) interface{}

	// Visit a parse tree produced by ASNParser#alternativeTypeList.
	VisitAlternativeTypeList(ctx *AlternativeTypeListContext) interface{}

	// Visit a parse tree produced by ASNParser#namedType.
	VisitNamedType(ctx *NamedTypeContext) interface{}

	// Visit a parse tree produced by ASNParser#enumeratedType.
	VisitEnumeratedType(ctx *EnumeratedTypeContext) interface{}

	// Visit a parse tree produced by ASNParser#enumerations.
	VisitEnumerations(ctx *EnumerationsContext) interface{}

	// Visit a parse tree produced by ASNParser#rootEnumeration.
	VisitRootEnumeration(ctx *RootEnumerationContext) interface{}

	// Visit a parse tree produced by ASNParser#enumeration.
	VisitEnumeration(ctx *EnumerationContext) interface{}

	// Visit a parse tree produced by ASNParser#enumerationItem.
	VisitEnumerationItem(ctx *EnumerationItemContext) interface{}

	// Visit a parse tree produced by ASNParser#namedNumber.
	VisitNamedNumber(ctx *NamedNumberContext) interface{}

	// Visit a parse tree produced by ASNParser#definedValue.
	VisitDefinedValue(ctx *DefinedValueContext) interface{}

	// Visit a parse tree produced by ASNParser#parameterizedValue.
	VisitParameterizedValue(ctx *ParameterizedValueContext) interface{}

	// Visit a parse tree produced by ASNParser#simpleDefinedValue.
	VisitSimpleDefinedValue(ctx *SimpleDefinedValueContext) interface{}

	// Visit a parse tree produced by ASNParser#actualParameterList.
	VisitActualParameterList(ctx *ActualParameterListContext) interface{}

	// Visit a parse tree produced by ASNParser#actualParameter.
	VisitActualParameter(ctx *ActualParameterContext) interface{}

	// Visit a parse tree produced by ASNParser#exceptionSpec.
	VisitExceptionSpec(ctx *ExceptionSpecContext) interface{}

	// Visit a parse tree produced by ASNParser#exceptionIdentification.
	VisitExceptionIdentification(ctx *ExceptionIdentificationContext) interface{}

	// Visit a parse tree produced by ASNParser#additionalEnumeration.
	VisitAdditionalEnumeration(ctx *AdditionalEnumerationContext) interface{}

	// Visit a parse tree produced by ASNParser#integerType.
	VisitIntegerType(ctx *IntegerTypeContext) interface{}

	// Visit a parse tree produced by ASNParser#namedNumberList.
	VisitNamedNumberList(ctx *NamedNumberListContext) interface{}

	// Visit a parse tree produced by ASNParser#objectidentifiertype.
	VisitObjectidentifiertype(ctx *ObjectidentifiertypeContext) interface{}

	// Visit a parse tree produced by ASNParser#componentRelationConstraint.
	VisitComponentRelationConstraint(ctx *ComponentRelationConstraintContext) interface{}

	// Visit a parse tree produced by ASNParser#atNotation.
	VisitAtNotation(ctx *AtNotationContext) interface{}

	// Visit a parse tree produced by ASNParser#level.
	VisitLevel(ctx *LevelContext) interface{}

	// Visit a parse tree produced by ASNParser#componentIdList.
	VisitComponentIdList(ctx *ComponentIdListContext) interface{}

	// Visit a parse tree produced by ASNParser#octetStringType.
	VisitOctetStringType(ctx *OctetStringTypeContext) interface{}

	// Visit a parse tree produced by ASNParser#bitStringType.
	VisitBitStringType(ctx *BitStringTypeContext) interface{}

	// Visit a parse tree produced by ASNParser#namedBitList.
	VisitNamedBitList(ctx *NamedBitListContext) interface{}

	// Visit a parse tree produced by ASNParser#namedBit.
	VisitNamedBit(ctx *NamedBitContext) interface{}

	// Visit a parse tree produced by ASNParser#booleanValue.
	VisitBooleanValue(ctx *BooleanValueContext) interface{}
}
