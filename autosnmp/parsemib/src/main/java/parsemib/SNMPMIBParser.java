// Generated from ./SNMPMIB.g4 by ANTLR 4.7.1
package parsemib;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class SNMPMIBParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.7.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, T__16=17, 
		T__17=18, T__18=19, T__19=20, T__20=21, T__21=22, T__22=23, T__23=24, 
		T__24=25, T__25=26, T__26=27, T__27=28, T__28=29, T__29=30, T__30=31, 
		T__31=32, T__32=33, T__33=34, T__34=35, T__35=36, T__36=37, T__37=38, 
		COMMENT=39, NUMBER=40, WS=41, LINE_COMMENT=42, BSTRING=43, HSTRING=44, 
		CSTRING=45, APOSTROPHE=46, IDENTIFIER=47;
	public static final int
		RULE_moduleDefinition = 0, RULE_tagDefault = 1, RULE_extensionDefault = 2, 
		RULE_moduleBody = 3, RULE_exports = 4, RULE_symbolsExported = 5, RULE_imports = 6, 
		RULE_symbolsImported = 7, RULE_symbolsFromModuleList = 8, RULE_symbolsFromModule = 9, 
		RULE_globalModuleReference = 10, RULE_symbolList = 11, RULE_symbol = 12, 
		RULE_assignmentList = 13, RULE_assignment = 14, RULE_typeAssignment = 15, 
		RULE_valueAssignment = 16, RULE_texttype = 17, RULE_asnType = 18, RULE_builtinType = 19, 
		RULE_valuetype = 20, RULE_modulevalueType = 21, RULE_objectvalueType = 22, 
		RULE_objectidentifiervaluetype = 23, RULE_valuedesctype = 24, RULE_syntaxdesctype = 25, 
		RULE_indexdesctype = 26, RULE_unknowndesctype = 27, RULE_octetStringType = 28, 
		RULE_objectidentifiertype = 29, RULE_sequencetype = 30, RULE_typedefine = 31, 
		RULE_integerType = 32, RULE_namedNumberList = 33, RULE_namedNumber = 34, 
		RULE_signedNumber = 35, RULE_sizeConstraint = 36, RULE_constraint = 37, 
		RULE_rangeconstraint = 38;
	public static final String[] ruleNames = {
		"moduleDefinition", "tagDefault", "extensionDefault", "moduleBody", "exports", 
		"symbolsExported", "imports", "symbolsImported", "symbolsFromModuleList", 
		"symbolsFromModule", "globalModuleReference", "symbolList", "symbol", 
		"assignmentList", "assignment", "typeAssignment", "valueAssignment", "texttype", 
		"asnType", "builtinType", "valuetype", "modulevalueType", "objectvalueType", 
		"objectidentifiervaluetype", "valuedesctype", "syntaxdesctype", "indexdesctype", 
		"unknowndesctype", "octetStringType", "objectidentifiertype", "sequencetype", 
		"typedefine", "integerType", "namedNumberList", "namedNumber", "signedNumber", 
		"sizeConstraint", "constraint", "rangeconstraint"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "'{'", "'('", "')'", "'}'", "'DEFINITIONS'", "'::='", "'BEGIN'", 
		"'END'", "'EXPLICIT'", "'IMPLICIT'", "'AUTOMATIC'", "'TAGS'", "'EXTENSIBILITY'", 
		"'IMPLIED'", "'EXPORTS'", "';'", "'ALL'", "'IMPORTS'", "'FROM'", "','", 
		"'MODULE-IDENTITY'", "'OBJECT-TYPE'", "'OBJECT'", "'IDENTIFIER'", "'SYNTAX'", 
		"'SEQUENCE'", "'OF'", "'INDEX'", "'OCTET'", "'STRING'", "'INTEGER'", "'INTEGER32'", 
		"'Integer32'", "'Integer'", "'|'", "'-'", "'SIZE'", "'..'", "'--'", null, 
		null, null, null, null, null, "'''"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, null, null, null, null, null, null, null, null, null, null, null, 
		null, null, null, null, null, null, null, null, null, null, null, null, 
		null, null, null, null, null, null, null, null, null, null, null, null, 
		null, null, null, "COMMENT", "NUMBER", "WS", "LINE_COMMENT", "BSTRING", 
		"HSTRING", "CSTRING", "APOSTROPHE", "IDENTIFIER"
	};
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "SNMPMIB.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public SNMPMIBParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}
	public static class ModuleDefinitionContext extends ParserRuleContext {
		public List<TerminalNode> IDENTIFIER() { return getTokens(SNMPMIBParser.IDENTIFIER); }
		public TerminalNode IDENTIFIER(int i) {
			return getToken(SNMPMIBParser.IDENTIFIER, i);
		}
		public TagDefaultContext tagDefault() {
			return getRuleContext(TagDefaultContext.class,0);
		}
		public ExtensionDefaultContext extensionDefault() {
			return getRuleContext(ExtensionDefaultContext.class,0);
		}
		public ModuleBodyContext moduleBody() {
			return getRuleContext(ModuleBodyContext.class,0);
		}
		public List<TerminalNode> NUMBER() { return getTokens(SNMPMIBParser.NUMBER); }
		public TerminalNode NUMBER(int i) {
			return getToken(SNMPMIBParser.NUMBER, i);
		}
		public ModuleDefinitionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_moduleDefinition; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).enterModuleDefinition(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).exitModuleDefinition(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof SNMPMIBVisitor ) return ((SNMPMIBVisitor<? extends T>)visitor).visitModuleDefinition(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ModuleDefinitionContext moduleDefinition() throws RecognitionException {
		ModuleDefinitionContext _localctx = new ModuleDefinitionContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_moduleDefinition);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(78);
			match(IDENTIFIER);
			setState(90);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__0) {
				{
				setState(79);
				match(T__0);
				setState(86);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==IDENTIFIER) {
					{
					{
					setState(80);
					match(IDENTIFIER);
					setState(81);
					match(T__1);
					setState(82);
					match(NUMBER);
					setState(83);
					match(T__2);
					}
					}
					setState(88);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(89);
				match(T__3);
				}
			}

			setState(92);
			match(T__4);
			setState(93);
			tagDefault();
			setState(94);
			extensionDefault();
			setState(95);
			match(T__5);
			setState(96);
			match(T__6);
			setState(97);
			moduleBody();
			setState(98);
			match(T__7);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TagDefaultContext extends ParserRuleContext {
		public TagDefaultContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tagDefault; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).enterTagDefault(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).exitTagDefault(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof SNMPMIBVisitor ) return ((SNMPMIBVisitor<? extends T>)visitor).visitTagDefault(this);
			else return visitor.visitChildren(this);
		}
	}

	public final TagDefaultContext tagDefault() throws RecognitionException {
		TagDefaultContext _localctx = new TagDefaultContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_tagDefault);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(102);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__8) | (1L << T__9) | (1L << T__10))) != 0)) {
				{
				setState(100);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__8) | (1L << T__9) | (1L << T__10))) != 0)) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(101);
				match(T__11);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExtensionDefaultContext extends ParserRuleContext {
		public ExtensionDefaultContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_extensionDefault; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).enterExtensionDefault(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).exitExtensionDefault(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof SNMPMIBVisitor ) return ((SNMPMIBVisitor<? extends T>)visitor).visitExtensionDefault(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ExtensionDefaultContext extensionDefault() throws RecognitionException {
		ExtensionDefaultContext _localctx = new ExtensionDefaultContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_extensionDefault);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(106);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__12) {
				{
				setState(104);
				match(T__12);
				setState(105);
				match(T__13);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ModuleBodyContext extends ParserRuleContext {
		public ExportsContext exports() {
			return getRuleContext(ExportsContext.class,0);
		}
		public ImportsContext imports() {
			return getRuleContext(ImportsContext.class,0);
		}
		public AssignmentListContext assignmentList() {
			return getRuleContext(AssignmentListContext.class,0);
		}
		public ModuleBodyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_moduleBody; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).enterModuleBody(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).exitModuleBody(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof SNMPMIBVisitor ) return ((SNMPMIBVisitor<? extends T>)visitor).visitModuleBody(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ModuleBodyContext moduleBody() throws RecognitionException {
		ModuleBodyContext _localctx = new ModuleBodyContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_moduleBody);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(112);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__14) | (1L << T__17) | (1L << IDENTIFIER))) != 0)) {
				{
				setState(108);
				exports();
				setState(109);
				imports();
				setState(110);
				assignmentList();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExportsContext extends ParserRuleContext {
		public SymbolsExportedContext symbolsExported() {
			return getRuleContext(SymbolsExportedContext.class,0);
		}
		public ExportsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_exports; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).enterExports(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).exitExports(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof SNMPMIBVisitor ) return ((SNMPMIBVisitor<? extends T>)visitor).visitExports(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ExportsContext exports() throws RecognitionException {
		ExportsContext _localctx = new ExportsContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_exports);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(121);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				{
				setState(114);
				match(T__14);
				setState(115);
				symbolsExported();
				setState(116);
				match(T__15);
				}
				break;
			case 2:
				{
				setState(118);
				match(T__14);
				setState(119);
				match(T__16);
				setState(120);
				match(T__15);
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SymbolsExportedContext extends ParserRuleContext {
		public SymbolListContext symbolList() {
			return getRuleContext(SymbolListContext.class,0);
		}
		public SymbolsExportedContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_symbolsExported; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).enterSymbolsExported(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).exitSymbolsExported(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof SNMPMIBVisitor ) return ((SNMPMIBVisitor<? extends T>)visitor).visitSymbolsExported(this);
			else return visitor.visitChildren(this);
		}
	}

	public final SymbolsExportedContext symbolsExported() throws RecognitionException {
		SymbolsExportedContext _localctx = new SymbolsExportedContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_symbolsExported);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(124);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__25) | (1L << T__28) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << IDENTIFIER))) != 0)) {
				{
				setState(123);
				symbolList();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ImportsContext extends ParserRuleContext {
		public SymbolsImportedContext symbolsImported() {
			return getRuleContext(SymbolsImportedContext.class,0);
		}
		public ImportsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_imports; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).enterImports(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).exitImports(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof SNMPMIBVisitor ) return ((SNMPMIBVisitor<? extends T>)visitor).visitImports(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ImportsContext imports() throws RecognitionException {
		ImportsContext _localctx = new ImportsContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_imports);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(130);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__17) {
				{
				setState(126);
				match(T__17);
				setState(127);
				symbolsImported();
				setState(128);
				match(T__15);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SymbolsImportedContext extends ParserRuleContext {
		public SymbolsFromModuleListContext symbolsFromModuleList() {
			return getRuleContext(SymbolsFromModuleListContext.class,0);
		}
		public SymbolsImportedContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_symbolsImported; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).enterSymbolsImported(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).exitSymbolsImported(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof SNMPMIBVisitor ) return ((SNMPMIBVisitor<? extends T>)visitor).visitSymbolsImported(this);
			else return visitor.visitChildren(this);
		}
	}

	public final SymbolsImportedContext symbolsImported() throws RecognitionException {
		SymbolsImportedContext _localctx = new SymbolsImportedContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_symbolsImported);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(133);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__25) | (1L << T__28) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << IDENTIFIER))) != 0)) {
				{
				setState(132);
				symbolsFromModuleList();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SymbolsFromModuleListContext extends ParserRuleContext {
		public List<SymbolsFromModuleContext> symbolsFromModule() {
			return getRuleContexts(SymbolsFromModuleContext.class);
		}
		public SymbolsFromModuleContext symbolsFromModule(int i) {
			return getRuleContext(SymbolsFromModuleContext.class,i);
		}
		public SymbolsFromModuleListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_symbolsFromModuleList; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).enterSymbolsFromModuleList(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).exitSymbolsFromModuleList(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof SNMPMIBVisitor ) return ((SNMPMIBVisitor<? extends T>)visitor).visitSymbolsFromModuleList(this);
			else return visitor.visitChildren(this);
		}
	}

	public final SymbolsFromModuleListContext symbolsFromModuleList() throws RecognitionException {
		SymbolsFromModuleListContext _localctx = new SymbolsFromModuleListContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_symbolsFromModuleList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(135);
			symbolsFromModule();
			}
			setState(139);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << T__25) | (1L << T__28) | (1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33) | (1L << IDENTIFIER))) != 0)) {
				{
				{
				setState(136);
				symbolsFromModule();
				}
				}
				setState(141);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SymbolsFromModuleContext extends ParserRuleContext {
		public SymbolListContext symbolList() {
			return getRuleContext(SymbolListContext.class,0);
		}
		public GlobalModuleReferenceContext globalModuleReference() {
			return getRuleContext(GlobalModuleReferenceContext.class,0);
		}
		public SymbolsFromModuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_symbolsFromModule; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).enterSymbolsFromModule(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).exitSymbolsFromModule(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof SNMPMIBVisitor ) return ((SNMPMIBVisitor<? extends T>)visitor).visitSymbolsFromModule(this);
			else return visitor.visitChildren(this);
		}
	}

	public final SymbolsFromModuleContext symbolsFromModule() throws RecognitionException {
		SymbolsFromModuleContext _localctx = new SymbolsFromModuleContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_symbolsFromModule);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(142);
			symbolList();
			setState(143);
			match(T__18);
			setState(144);
			globalModuleReference();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class GlobalModuleReferenceContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(SNMPMIBParser.IDENTIFIER, 0); }
		public GlobalModuleReferenceContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_globalModuleReference; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).enterGlobalModuleReference(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).exitGlobalModuleReference(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof SNMPMIBVisitor ) return ((SNMPMIBVisitor<? extends T>)visitor).visitGlobalModuleReference(this);
			else return visitor.visitChildren(this);
		}
	}

	public final GlobalModuleReferenceContext globalModuleReference() throws RecognitionException {
		GlobalModuleReferenceContext _localctx = new GlobalModuleReferenceContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_globalModuleReference);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(146);
			match(IDENTIFIER);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SymbolListContext extends ParserRuleContext {
		public List<SymbolContext> symbol() {
			return getRuleContexts(SymbolContext.class);
		}
		public SymbolContext symbol(int i) {
			return getRuleContext(SymbolContext.class,i);
		}
		public SymbolListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_symbolList; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).enterSymbolList(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).exitSymbolList(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof SNMPMIBVisitor ) return ((SNMPMIBVisitor<? extends T>)visitor).visitSymbolList(this);
			else return visitor.visitChildren(this);
		}
	}

	public final SymbolListContext symbolList() throws RecognitionException {
		SymbolListContext _localctx = new SymbolListContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_symbolList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(148);
			symbol();
			}
			setState(153);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__19) {
				{
				{
				setState(149);
				match(T__19);
				setState(150);
				symbol();
				}
				}
				setState(155);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SymbolContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(SNMPMIBParser.IDENTIFIER, 0); }
		public ValuetypeContext valuetype() {
			return getRuleContext(ValuetypeContext.class,0);
		}
		public AsnTypeContext asnType() {
			return getRuleContext(AsnTypeContext.class,0);
		}
		public SymbolContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_symbol; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).enterSymbol(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).exitSymbol(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof SNMPMIBVisitor ) return ((SNMPMIBVisitor<? extends T>)visitor).visitSymbol(this);
			else return visitor.visitChildren(this);
		}
	}

	public final SymbolContext symbol() throws RecognitionException {
		SymbolContext _localctx = new SymbolContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_symbol);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(159);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
			case 1:
				{
				setState(156);
				match(IDENTIFIER);
				}
				break;
			case 2:
				{
				setState(157);
				valuetype();
				}
				break;
			case 3:
				{
				setState(158);
				asnType();
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AssignmentListContext extends ParserRuleContext {
		public List<AssignmentContext> assignment() {
			return getRuleContexts(AssignmentContext.class);
		}
		public AssignmentContext assignment(int i) {
			return getRuleContext(AssignmentContext.class,i);
		}
		public AssignmentListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignmentList; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).enterAssignmentList(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).exitAssignmentList(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof SNMPMIBVisitor ) return ((SNMPMIBVisitor<? extends T>)visitor).visitAssignmentList(this);
			else return visitor.visitChildren(this);
		}
	}

	public final AssignmentListContext assignmentList() throws RecognitionException {
		AssignmentListContext _localctx = new AssignmentListContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_assignmentList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(161);
			assignment();
			}
			setState(165);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==IDENTIFIER) {
				{
				{
				setState(162);
				assignment();
				}
				}
				setState(167);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AssignmentContext extends ParserRuleContext {
		public ValueAssignmentContext valueAssignment() {
			return getRuleContext(ValueAssignmentContext.class,0);
		}
		public TypeAssignmentContext typeAssignment() {
			return getRuleContext(TypeAssignmentContext.class,0);
		}
		public AssignmentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignment; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).enterAssignment(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).exitAssignment(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof SNMPMIBVisitor ) return ((SNMPMIBVisitor<? extends T>)visitor).visitAssignment(this);
			else return visitor.visitChildren(this);
		}
	}

	public final AssignmentContext assignment() throws RecognitionException {
		AssignmentContext _localctx = new AssignmentContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_assignment);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(170);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				{
				setState(168);
				valueAssignment();
				}
				break;
			case 2:
				{
				setState(169);
				typeAssignment();
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypeAssignmentContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(SNMPMIBParser.IDENTIFIER, 0); }
		public AsnTypeContext asnType() {
			return getRuleContext(AsnTypeContext.class,0);
		}
		public TexttypeContext texttype() {
			return getRuleContext(TexttypeContext.class,0);
		}
		public TypeAssignmentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeAssignment; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).enterTypeAssignment(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).exitTypeAssignment(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof SNMPMIBVisitor ) return ((SNMPMIBVisitor<? extends T>)visitor).visitTypeAssignment(this);
			else return visitor.visitChildren(this);
		}
	}

	public final TypeAssignmentContext typeAssignment() throws RecognitionException {
		TypeAssignmentContext _localctx = new TypeAssignmentContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_typeAssignment);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(172);
			match(IDENTIFIER);
			setState(173);
			match(T__5);
			setState(176);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
			case 1:
				{
				setState(174);
				asnType();
				}
				break;
			case 2:
				{
				setState(175);
				texttype();
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ValueAssignmentContext extends ParserRuleContext {
		public List<TerminalNode> IDENTIFIER() { return getTokens(SNMPMIBParser.IDENTIFIER); }
		public TerminalNode IDENTIFIER(int i) {
			return getToken(SNMPMIBParser.IDENTIFIER, i);
		}
		public ValuetypeContext valuetype() {
			return getRuleContext(ValuetypeContext.class,0);
		}
		public TerminalNode NUMBER() { return getToken(SNMPMIBParser.NUMBER, 0); }
		public List<ValuedesctypeContext> valuedesctype() {
			return getRuleContexts(ValuedesctypeContext.class);
		}
		public ValuedesctypeContext valuedesctype(int i) {
			return getRuleContext(ValuedesctypeContext.class,i);
		}
		public ValueAssignmentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_valueAssignment; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).enterValueAssignment(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).exitValueAssignment(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof SNMPMIBVisitor ) return ((SNMPMIBVisitor<? extends T>)visitor).visitValueAssignment(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ValueAssignmentContext valueAssignment() throws RecognitionException {
		ValueAssignmentContext _localctx = new ValueAssignmentContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_valueAssignment);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(178);
			match(IDENTIFIER);
			setState(179);
			valuetype();
			setState(183);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__22) | (1L << T__24) | (1L << T__27) | (1L << IDENTIFIER))) != 0)) {
				{
				{
				setState(180);
				valuedesctype();
				}
				}
				setState(185);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(186);
			match(T__5);
			setState(187);
			match(T__0);
			setState(188);
			match(IDENTIFIER);
			setState(189);
			match(NUMBER);
			setState(190);
			match(T__3);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TexttypeContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(SNMPMIBParser.IDENTIFIER, 0); }
		public List<ValuedesctypeContext> valuedesctype() {
			return getRuleContexts(ValuedesctypeContext.class);
		}
		public ValuedesctypeContext valuedesctype(int i) {
			return getRuleContext(ValuedesctypeContext.class,i);
		}
		public TexttypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_texttype; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).enterTexttype(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).exitTexttype(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof SNMPMIBVisitor ) return ((SNMPMIBVisitor<? extends T>)visitor).visitTexttype(this);
			else return visitor.visitChildren(this);
		}
	}

	public final TexttypeContext texttype() throws RecognitionException {
		TexttypeContext _localctx = new TexttypeContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_texttype);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(192);
			match(IDENTIFIER);
			setState(196);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,16,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(193);
					valuedesctype();
					}
					} 
				}
				setState(198);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,16,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AsnTypeContext extends ParserRuleContext {
		public BuiltinTypeContext builtinType() {
			return getRuleContext(BuiltinTypeContext.class,0);
		}
		public TerminalNode IDENTIFIER() { return getToken(SNMPMIBParser.IDENTIFIER, 0); }
		public List<ConstraintContext> constraint() {
			return getRuleContexts(ConstraintContext.class);
		}
		public ConstraintContext constraint(int i) {
			return getRuleContext(ConstraintContext.class,i);
		}
		public List<SizeConstraintContext> sizeConstraint() {
			return getRuleContexts(SizeConstraintContext.class);
		}
		public SizeConstraintContext sizeConstraint(int i) {
			return getRuleContext(SizeConstraintContext.class,i);
		}
		public AsnTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_asnType; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).enterAsnType(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).exitAsnType(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof SNMPMIBVisitor ) return ((SNMPMIBVisitor<? extends T>)visitor).visitAsnType(this);
			else return visitor.visitChildren(this);
		}
	}

	public final AsnTypeContext asnType() throws RecognitionException {
		AsnTypeContext _localctx = new AsnTypeContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_asnType);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(201);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__22:
			case T__25:
			case T__28:
			case T__30:
			case T__31:
			case T__32:
			case T__33:
				{
				setState(199);
				builtinType();
				}
				break;
			case IDENTIFIER:
				{
				setState(200);
				match(IDENTIFIER);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(207);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__1) {
				{
				setState(205);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,18,_ctx) ) {
				case 1:
					{
					setState(203);
					constraint();
					}
					break;
				case 2:
					{
					setState(204);
					sizeConstraint();
					}
					break;
				}
				}
				setState(209);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class BuiltinTypeContext extends ParserRuleContext {
		public OctetStringTypeContext octetStringType() {
			return getRuleContext(OctetStringTypeContext.class,0);
		}
		public IntegerTypeContext integerType() {
			return getRuleContext(IntegerTypeContext.class,0);
		}
		public ObjectidentifiertypeContext objectidentifiertype() {
			return getRuleContext(ObjectidentifiertypeContext.class,0);
		}
		public SequencetypeContext sequencetype() {
			return getRuleContext(SequencetypeContext.class,0);
		}
		public BuiltinTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_builtinType; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).enterBuiltinType(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).exitBuiltinType(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof SNMPMIBVisitor ) return ((SNMPMIBVisitor<? extends T>)visitor).visitBuiltinType(this);
			else return visitor.visitChildren(this);
		}
	}

	public final BuiltinTypeContext builtinType() throws RecognitionException {
		BuiltinTypeContext _localctx = new BuiltinTypeContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_builtinType);
		try {
			setState(214);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__28:
				enterOuterAlt(_localctx, 1);
				{
				setState(210);
				octetStringType();
				}
				break;
			case T__30:
			case T__31:
			case T__32:
			case T__33:
				enterOuterAlt(_localctx, 2);
				{
				setState(211);
				integerType();
				}
				break;
			case T__22:
				enterOuterAlt(_localctx, 3);
				{
				setState(212);
				objectidentifiertype();
				}
				break;
			case T__25:
				enterOuterAlt(_localctx, 4);
				{
				setState(213);
				sequencetype();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ValuetypeContext extends ParserRuleContext {
		public ObjectidentifiervaluetypeContext objectidentifiervaluetype() {
			return getRuleContext(ObjectidentifiervaluetypeContext.class,0);
		}
		public ObjectvalueTypeContext objectvalueType() {
			return getRuleContext(ObjectvalueTypeContext.class,0);
		}
		public ModulevalueTypeContext modulevalueType() {
			return getRuleContext(ModulevalueTypeContext.class,0);
		}
		public TerminalNode IDENTIFIER() { return getToken(SNMPMIBParser.IDENTIFIER, 0); }
		public ValuetypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_valuetype; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).enterValuetype(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).exitValuetype(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof SNMPMIBVisitor ) return ((SNMPMIBVisitor<? extends T>)visitor).visitValuetype(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ValuetypeContext valuetype() throws RecognitionException {
		ValuetypeContext _localctx = new ValuetypeContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_valuetype);
		try {
			setState(220);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__22:
				enterOuterAlt(_localctx, 1);
				{
				setState(216);
				objectidentifiervaluetype();
				}
				break;
			case T__21:
				enterOuterAlt(_localctx, 2);
				{
				setState(217);
				objectvalueType();
				}
				break;
			case T__20:
				enterOuterAlt(_localctx, 3);
				{
				setState(218);
				modulevalueType();
				}
				break;
			case IDENTIFIER:
				enterOuterAlt(_localctx, 4);
				{
				setState(219);
				match(IDENTIFIER);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ModulevalueTypeContext extends ParserRuleContext {
		public ModulevalueTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_modulevalueType; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).enterModulevalueType(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).exitModulevalueType(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof SNMPMIBVisitor ) return ((SNMPMIBVisitor<? extends T>)visitor).visitModulevalueType(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ModulevalueTypeContext modulevalueType() throws RecognitionException {
		ModulevalueTypeContext _localctx = new ModulevalueTypeContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_modulevalueType);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(222);
			match(T__20);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ObjectvalueTypeContext extends ParserRuleContext {
		public ObjectvalueTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_objectvalueType; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).enterObjectvalueType(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).exitObjectvalueType(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof SNMPMIBVisitor ) return ((SNMPMIBVisitor<? extends T>)visitor).visitObjectvalueType(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ObjectvalueTypeContext objectvalueType() throws RecognitionException {
		ObjectvalueTypeContext _localctx = new ObjectvalueTypeContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_objectvalueType);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(224);
			match(T__21);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ObjectidentifiervaluetypeContext extends ParserRuleContext {
		public ObjectidentifiervaluetypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_objectidentifiervaluetype; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).enterObjectidentifiervaluetype(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).exitObjectidentifiervaluetype(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof SNMPMIBVisitor ) return ((SNMPMIBVisitor<? extends T>)visitor).visitObjectidentifiervaluetype(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ObjectidentifiervaluetypeContext objectidentifiervaluetype() throws RecognitionException {
		ObjectidentifiervaluetypeContext _localctx = new ObjectidentifiervaluetypeContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_objectidentifiervaluetype);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(226);
			match(T__22);
			setState(227);
			match(T__23);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ValuedesctypeContext extends ParserRuleContext {
		public SyntaxdesctypeContext syntaxdesctype() {
			return getRuleContext(SyntaxdesctypeContext.class,0);
		}
		public IndexdesctypeContext indexdesctype() {
			return getRuleContext(IndexdesctypeContext.class,0);
		}
		public UnknowndesctypeContext unknowndesctype() {
			return getRuleContext(UnknowndesctypeContext.class,0);
		}
		public ValuedesctypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_valuedesctype; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).enterValuedesctype(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).exitValuedesctype(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof SNMPMIBVisitor ) return ((SNMPMIBVisitor<? extends T>)visitor).visitValuedesctype(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ValuedesctypeContext valuedesctype() throws RecognitionException {
		ValuedesctypeContext _localctx = new ValuedesctypeContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_valuedesctype);
		try {
			setState(232);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,22,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(229);
				syntaxdesctype();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(230);
				indexdesctype();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(231);
				unknowndesctype();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SyntaxdesctypeContext extends ParserRuleContext {
		public AsnTypeContext asnType() {
			return getRuleContext(AsnTypeContext.class,0);
		}
		public TerminalNode IDENTIFIER() { return getToken(SNMPMIBParser.IDENTIFIER, 0); }
		public SyntaxdesctypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_syntaxdesctype; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).enterSyntaxdesctype(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).exitSyntaxdesctype(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof SNMPMIBVisitor ) return ((SNMPMIBVisitor<? extends T>)visitor).visitSyntaxdesctype(this);
			else return visitor.visitChildren(this);
		}
	}

	public final SyntaxdesctypeContext syntaxdesctype() throws RecognitionException {
		SyntaxdesctypeContext _localctx = new SyntaxdesctypeContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_syntaxdesctype);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(234);
			match(T__24);
			setState(239);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,23,_ctx) ) {
			case 1:
				{
				setState(235);
				asnType();
				}
				break;
			case 2:
				{
				{
				setState(236);
				match(T__25);
				setState(237);
				match(T__26);
				setState(238);
				match(IDENTIFIER);
				}
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class IndexdesctypeContext extends ParserRuleContext {
		public List<TerminalNode> IDENTIFIER() { return getTokens(SNMPMIBParser.IDENTIFIER); }
		public TerminalNode IDENTIFIER(int i) {
			return getToken(SNMPMIBParser.IDENTIFIER, i);
		}
		public IndexdesctypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_indexdesctype; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).enterIndexdesctype(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).exitIndexdesctype(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof SNMPMIBVisitor ) return ((SNMPMIBVisitor<? extends T>)visitor).visitIndexdesctype(this);
			else return visitor.visitChildren(this);
		}
	}

	public final IndexdesctypeContext indexdesctype() throws RecognitionException {
		IndexdesctypeContext _localctx = new IndexdesctypeContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_indexdesctype);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(241);
			_la = _input.LA(1);
			if ( !(_la==T__27 || _la==IDENTIFIER) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(243);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==IDENTIFIER) {
				{
				setState(242);
				match(IDENTIFIER);
				}
			}

			setState(245);
			match(T__0);
			setState(246);
			match(IDENTIFIER);
			setState(251);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__19) {
				{
				{
				setState(247);
				match(T__19);
				setState(248);
				match(IDENTIFIER);
				}
				}
				setState(253);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(254);
			match(T__3);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class UnknowndesctypeContext extends ParserRuleContext {
		public List<TerminalNode> IDENTIFIER() { return getTokens(SNMPMIBParser.IDENTIFIER); }
		public TerminalNode IDENTIFIER(int i) {
			return getToken(SNMPMIBParser.IDENTIFIER, i);
		}
		public TerminalNode CSTRING() { return getToken(SNMPMIBParser.CSTRING, 0); }
		public UnknowndesctypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_unknowndesctype; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).enterUnknowndesctype(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).exitUnknowndesctype(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof SNMPMIBVisitor ) return ((SNMPMIBVisitor<? extends T>)visitor).visitUnknowndesctype(this);
			else return visitor.visitChildren(this);
		}
	}

	public final UnknowndesctypeContext unknowndesctype() throws RecognitionException {
		UnknowndesctypeContext _localctx = new UnknowndesctypeContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_unknowndesctype);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(256);
			_la = _input.LA(1);
			if ( !(_la==T__22 || _la==IDENTIFIER) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(257);
			_la = _input.LA(1);
			if ( !(_la==CSTRING || _la==IDENTIFIER) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class OctetStringTypeContext extends ParserRuleContext {
		public OctetStringTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_octetStringType; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).enterOctetStringType(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).exitOctetStringType(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof SNMPMIBVisitor ) return ((SNMPMIBVisitor<? extends T>)visitor).visitOctetStringType(this);
			else return visitor.visitChildren(this);
		}
	}

	public final OctetStringTypeContext octetStringType() throws RecognitionException {
		OctetStringTypeContext _localctx = new OctetStringTypeContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_octetStringType);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(259);
			match(T__28);
			setState(260);
			match(T__29);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ObjectidentifiertypeContext extends ParserRuleContext {
		public ObjectidentifiertypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_objectidentifiertype; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).enterObjectidentifiertype(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).exitObjectidentifiertype(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof SNMPMIBVisitor ) return ((SNMPMIBVisitor<? extends T>)visitor).visitObjectidentifiertype(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ObjectidentifiertypeContext objectidentifiertype() throws RecognitionException {
		ObjectidentifiertypeContext _localctx = new ObjectidentifiertypeContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_objectidentifiertype);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(262);
			match(T__22);
			setState(263);
			match(T__23);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SequencetypeContext extends ParserRuleContext {
		public List<TypedefineContext> typedefine() {
			return getRuleContexts(TypedefineContext.class);
		}
		public TypedefineContext typedefine(int i) {
			return getRuleContext(TypedefineContext.class,i);
		}
		public SequencetypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sequencetype; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).enterSequencetype(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).exitSequencetype(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof SNMPMIBVisitor ) return ((SNMPMIBVisitor<? extends T>)visitor).visitSequencetype(this);
			else return visitor.visitChildren(this);
		}
	}

	public final SequencetypeContext sequencetype() throws RecognitionException {
		SequencetypeContext _localctx = new SequencetypeContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_sequencetype);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(265);
			match(T__25);
			setState(266);
			match(T__0);
			setState(267);
			typedefine();
			setState(272);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__19) {
				{
				{
				setState(268);
				match(T__19);
				setState(269);
				typedefine();
				}
				}
				setState(274);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(275);
			match(T__3);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypedefineContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(SNMPMIBParser.IDENTIFIER, 0); }
		public AsnTypeContext asnType() {
			return getRuleContext(AsnTypeContext.class,0);
		}
		public TypedefineContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typedefine; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).enterTypedefine(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).exitTypedefine(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof SNMPMIBVisitor ) return ((SNMPMIBVisitor<? extends T>)visitor).visitTypedefine(this);
			else return visitor.visitChildren(this);
		}
	}

	public final TypedefineContext typedefine() throws RecognitionException {
		TypedefineContext _localctx = new TypedefineContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_typedefine);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(277);
			match(IDENTIFIER);
			setState(278);
			asnType();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class IntegerTypeContext extends ParserRuleContext {
		public NamedNumberListContext namedNumberList() {
			return getRuleContext(NamedNumberListContext.class,0);
		}
		public List<TerminalNode> NUMBER() { return getTokens(SNMPMIBParser.NUMBER); }
		public TerminalNode NUMBER(int i) {
			return getToken(SNMPMIBParser.NUMBER, i);
		}
		public IntegerTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_integerType; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).enterIntegerType(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).exitIntegerType(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof SNMPMIBVisitor ) return ((SNMPMIBVisitor<? extends T>)visitor).visitIntegerType(this);
			else return visitor.visitChildren(this);
		}
	}

	public final IntegerTypeContext integerType() throws RecognitionException {
		IntegerTypeContext _localctx = new IntegerTypeContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_integerType);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(280);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__30) | (1L << T__31) | (1L << T__32) | (1L << T__33))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(294);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,28,_ctx) ) {
			case 1:
				{
				{
				setState(281);
				match(T__0);
				setState(282);
				namedNumberList();
				setState(283);
				match(T__3);
				}
				}
				break;
			case 2:
				{
				{
				setState(285);
				match(T__1);
				setState(286);
				match(NUMBER);
				setState(289); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(287);
					match(T__34);
					setState(288);
					match(NUMBER);
					}
					}
					setState(291); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==T__34 );
				setState(293);
				match(T__2);
				}
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class NamedNumberListContext extends ParserRuleContext {
		public List<NamedNumberContext> namedNumber() {
			return getRuleContexts(NamedNumberContext.class);
		}
		public NamedNumberContext namedNumber(int i) {
			return getRuleContext(NamedNumberContext.class,i);
		}
		public NamedNumberListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_namedNumberList; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).enterNamedNumberList(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).exitNamedNumberList(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof SNMPMIBVisitor ) return ((SNMPMIBVisitor<? extends T>)visitor).visitNamedNumberList(this);
			else return visitor.visitChildren(this);
		}
	}

	public final NamedNumberListContext namedNumberList() throws RecognitionException {
		NamedNumberListContext _localctx = new NamedNumberListContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_namedNumberList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(296);
			namedNumber();
			}
			setState(301);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__19) {
				{
				{
				setState(297);
				match(T__19);
				setState(298);
				namedNumber();
				}
				}
				setState(303);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class NamedNumberContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(SNMPMIBParser.IDENTIFIER, 0); }
		public SignedNumberContext signedNumber() {
			return getRuleContext(SignedNumberContext.class,0);
		}
		public NamedNumberContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_namedNumber; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).enterNamedNumber(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).exitNamedNumber(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof SNMPMIBVisitor ) return ((SNMPMIBVisitor<? extends T>)visitor).visitNamedNumber(this);
			else return visitor.visitChildren(this);
		}
	}

	public final NamedNumberContext namedNumber() throws RecognitionException {
		NamedNumberContext _localctx = new NamedNumberContext(_ctx, getState());
		enterRule(_localctx, 68, RULE_namedNumber);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(304);
			match(IDENTIFIER);
			setState(305);
			match(T__1);
			setState(306);
			signedNumber();
			setState(307);
			match(T__2);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SignedNumberContext extends ParserRuleContext {
		public TerminalNode NUMBER() { return getToken(SNMPMIBParser.NUMBER, 0); }
		public SignedNumberContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_signedNumber; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).enterSignedNumber(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).exitSignedNumber(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof SNMPMIBVisitor ) return ((SNMPMIBVisitor<? extends T>)visitor).visitSignedNumber(this);
			else return visitor.visitChildren(this);
		}
	}

	public final SignedNumberContext signedNumber() throws RecognitionException {
		SignedNumberContext _localctx = new SignedNumberContext(_ctx, getState());
		enterRule(_localctx, 70, RULE_signedNumber);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(310);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__35) {
				{
				setState(309);
				match(T__35);
				}
			}

			setState(312);
			match(NUMBER);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SizeConstraintContext extends ParserRuleContext {
		public ConstraintContext constraint() {
			return getRuleContext(ConstraintContext.class,0);
		}
		public SizeConstraintContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sizeConstraint; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).enterSizeConstraint(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).exitSizeConstraint(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof SNMPMIBVisitor ) return ((SNMPMIBVisitor<? extends T>)visitor).visitSizeConstraint(this);
			else return visitor.visitChildren(this);
		}
	}

	public final SizeConstraintContext sizeConstraint() throws RecognitionException {
		SizeConstraintContext _localctx = new SizeConstraintContext(_ctx, getState());
		enterRule(_localctx, 72, RULE_sizeConstraint);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(314);
			match(T__1);
			setState(315);
			match(T__36);
			setState(316);
			constraint();
			setState(317);
			match(T__2);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ConstraintContext extends ParserRuleContext {
		public RangeconstraintContext rangeconstraint() {
			return getRuleContext(RangeconstraintContext.class,0);
		}
		public TerminalNode NUMBER() { return getToken(SNMPMIBParser.NUMBER, 0); }
		public ConstraintContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_constraint; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).enterConstraint(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).exitConstraint(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof SNMPMIBVisitor ) return ((SNMPMIBVisitor<? extends T>)visitor).visitConstraint(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ConstraintContext constraint() throws RecognitionException {
		ConstraintContext _localctx = new ConstraintContext(_ctx, getState());
		enterRule(_localctx, 74, RULE_constraint);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(319);
			match(T__1);
			setState(322);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,31,_ctx) ) {
			case 1:
				{
				setState(320);
				rangeconstraint();
				}
				break;
			case 2:
				{
				setState(321);
				match(NUMBER);
				}
				break;
			}
			setState(324);
			match(T__2);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class RangeconstraintContext extends ParserRuleContext {
		public List<SignedNumberContext> signedNumber() {
			return getRuleContexts(SignedNumberContext.class);
		}
		public SignedNumberContext signedNumber(int i) {
			return getRuleContext(SignedNumberContext.class,i);
		}
		public RangeconstraintContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_rangeconstraint; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).enterRangeconstraint(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SNMPMIBListener ) ((SNMPMIBListener)listener).exitRangeconstraint(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof SNMPMIBVisitor ) return ((SNMPMIBVisitor<? extends T>)visitor).visitRangeconstraint(this);
			else return visitor.visitChildren(this);
		}
	}

	public final RangeconstraintContext rangeconstraint() throws RecognitionException {
		RangeconstraintContext _localctx = new RangeconstraintContext(_ctx, getState());
		enterRule(_localctx, 76, RULE_rangeconstraint);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(326);
			signedNumber();
			setState(327);
			match(T__37);
			setState(328);
			signedNumber();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\61\u014d\4\2\t\2"+
		"\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\3\2\3\2\3\2\3\2\3\2"+
		"\3\2\7\2W\n\2\f\2\16\2Z\13\2\3\2\5\2]\n\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2"+
		"\3\2\3\3\3\3\5\3i\n\3\3\4\3\4\5\4m\n\4\3\5\3\5\3\5\3\5\5\5s\n\5\3\6\3"+
		"\6\3\6\3\6\3\6\3\6\3\6\5\6|\n\6\3\7\5\7\177\n\7\3\b\3\b\3\b\3\b\5\b\u0085"+
		"\n\b\3\t\5\t\u0088\n\t\3\n\3\n\7\n\u008c\n\n\f\n\16\n\u008f\13\n\3\13"+
		"\3\13\3\13\3\13\3\f\3\f\3\r\3\r\3\r\7\r\u009a\n\r\f\r\16\r\u009d\13\r"+
		"\3\16\3\16\3\16\5\16\u00a2\n\16\3\17\3\17\7\17\u00a6\n\17\f\17\16\17\u00a9"+
		"\13\17\3\20\3\20\5\20\u00ad\n\20\3\21\3\21\3\21\3\21\5\21\u00b3\n\21\3"+
		"\22\3\22\3\22\7\22\u00b8\n\22\f\22\16\22\u00bb\13\22\3\22\3\22\3\22\3"+
		"\22\3\22\3\22\3\23\3\23\7\23\u00c5\n\23\f\23\16\23\u00c8\13\23\3\24\3"+
		"\24\5\24\u00cc\n\24\3\24\3\24\7\24\u00d0\n\24\f\24\16\24\u00d3\13\24\3"+
		"\25\3\25\3\25\3\25\5\25\u00d9\n\25\3\26\3\26\3\26\3\26\5\26\u00df\n\26"+
		"\3\27\3\27\3\30\3\30\3\31\3\31\3\31\3\32\3\32\3\32\5\32\u00eb\n\32\3\33"+
		"\3\33\3\33\3\33\3\33\5\33\u00f2\n\33\3\34\3\34\5\34\u00f6\n\34\3\34\3"+
		"\34\3\34\3\34\7\34\u00fc\n\34\f\34\16\34\u00ff\13\34\3\34\3\34\3\35\3"+
		"\35\3\35\3\36\3\36\3\36\3\37\3\37\3\37\3 \3 \3 \3 \3 \7 \u0111\n \f \16"+
		" \u0114\13 \3 \3 \3!\3!\3!\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\6\"\u0124"+
		"\n\"\r\"\16\"\u0125\3\"\5\"\u0129\n\"\3#\3#\3#\7#\u012e\n#\f#\16#\u0131"+
		"\13#\3$\3$\3$\3$\3$\3%\5%\u0139\n%\3%\3%\3&\3&\3&\3&\3&\3\'\3\'\3\'\5"+
		"\'\u0145\n\'\3\'\3\'\3(\3(\3(\3(\3(\2\2)\2\4\6\b\n\f\16\20\22\24\26\30"+
		"\32\34\36 \"$&(*,.\60\62\64\668:<>@BDFHJLN\2\7\3\2\13\r\4\2\36\36\61\61"+
		"\4\2\31\31\61\61\4\2//\61\61\3\2!$\2\u014d\2P\3\2\2\2\4h\3\2\2\2\6l\3"+
		"\2\2\2\br\3\2\2\2\n{\3\2\2\2\f~\3\2\2\2\16\u0084\3\2\2\2\20\u0087\3\2"+
		"\2\2\22\u0089\3\2\2\2\24\u0090\3\2\2\2\26\u0094\3\2\2\2\30\u0096\3\2\2"+
		"\2\32\u00a1\3\2\2\2\34\u00a3\3\2\2\2\36\u00ac\3\2\2\2 \u00ae\3\2\2\2\""+
		"\u00b4\3\2\2\2$\u00c2\3\2\2\2&\u00cb\3\2\2\2(\u00d8\3\2\2\2*\u00de\3\2"+
		"\2\2,\u00e0\3\2\2\2.\u00e2\3\2\2\2\60\u00e4\3\2\2\2\62\u00ea\3\2\2\2\64"+
		"\u00ec\3\2\2\2\66\u00f3\3\2\2\28\u0102\3\2\2\2:\u0105\3\2\2\2<\u0108\3"+
		"\2\2\2>\u010b\3\2\2\2@\u0117\3\2\2\2B\u011a\3\2\2\2D\u012a\3\2\2\2F\u0132"+
		"\3\2\2\2H\u0138\3\2\2\2J\u013c\3\2\2\2L\u0141\3\2\2\2N\u0148\3\2\2\2P"+
		"\\\7\61\2\2QX\7\3\2\2RS\7\61\2\2ST\7\4\2\2TU\7*\2\2UW\7\5\2\2VR\3\2\2"+
		"\2WZ\3\2\2\2XV\3\2\2\2XY\3\2\2\2Y[\3\2\2\2ZX\3\2\2\2[]\7\6\2\2\\Q\3\2"+
		"\2\2\\]\3\2\2\2]^\3\2\2\2^_\7\7\2\2_`\5\4\3\2`a\5\6\4\2ab\7\b\2\2bc\7"+
		"\t\2\2cd\5\b\5\2de\7\n\2\2e\3\3\2\2\2fg\t\2\2\2gi\7\16\2\2hf\3\2\2\2h"+
		"i\3\2\2\2i\5\3\2\2\2jk\7\17\2\2km\7\20\2\2lj\3\2\2\2lm\3\2\2\2m\7\3\2"+
		"\2\2no\5\n\6\2op\5\16\b\2pq\5\34\17\2qs\3\2\2\2rn\3\2\2\2rs\3\2\2\2s\t"+
		"\3\2\2\2tu\7\21\2\2uv\5\f\7\2vw\7\22\2\2w|\3\2\2\2xy\7\21\2\2yz\7\23\2"+
		"\2z|\7\22\2\2{t\3\2\2\2{x\3\2\2\2{|\3\2\2\2|\13\3\2\2\2}\177\5\30\r\2"+
		"~}\3\2\2\2~\177\3\2\2\2\177\r\3\2\2\2\u0080\u0081\7\24\2\2\u0081\u0082"+
		"\5\20\t\2\u0082\u0083\7\22\2\2\u0083\u0085\3\2\2\2\u0084\u0080\3\2\2\2"+
		"\u0084\u0085\3\2\2\2\u0085\17\3\2\2\2\u0086\u0088\5\22\n\2\u0087\u0086"+
		"\3\2\2\2\u0087\u0088\3\2\2\2\u0088\21\3\2\2\2\u0089\u008d\5\24\13\2\u008a"+
		"\u008c\5\24\13\2\u008b\u008a\3\2\2\2\u008c\u008f\3\2\2\2\u008d\u008b\3"+
		"\2\2\2\u008d\u008e\3\2\2\2\u008e\23\3\2\2\2\u008f\u008d\3\2\2\2\u0090"+
		"\u0091\5\30\r\2\u0091\u0092\7\25\2\2\u0092\u0093\5\26\f\2\u0093\25\3\2"+
		"\2\2\u0094\u0095\7\61\2\2\u0095\27\3\2\2\2\u0096\u009b\5\32\16\2\u0097"+
		"\u0098\7\26\2\2\u0098\u009a\5\32\16\2\u0099\u0097\3\2\2\2\u009a\u009d"+
		"\3\2\2\2\u009b\u0099\3\2\2\2\u009b\u009c\3\2\2\2\u009c\31\3\2\2\2\u009d"+
		"\u009b\3\2\2\2\u009e\u00a2\7\61\2\2\u009f\u00a2\5*\26\2\u00a0\u00a2\5"+
		"&\24\2\u00a1\u009e\3\2\2\2\u00a1\u009f\3\2\2\2\u00a1\u00a0\3\2\2\2\u00a2"+
		"\33\3\2\2\2\u00a3\u00a7\5\36\20\2\u00a4\u00a6\5\36\20\2\u00a5\u00a4\3"+
		"\2\2\2\u00a6\u00a9\3\2\2\2\u00a7\u00a5\3\2\2\2\u00a7\u00a8\3\2\2\2\u00a8"+
		"\35\3\2\2\2\u00a9\u00a7\3\2\2\2\u00aa\u00ad\5\"\22\2\u00ab\u00ad\5 \21"+
		"\2\u00ac\u00aa\3\2\2\2\u00ac\u00ab\3\2\2\2\u00ad\37\3\2\2\2\u00ae\u00af"+
		"\7\61\2\2\u00af\u00b2\7\b\2\2\u00b0\u00b3\5&\24\2\u00b1\u00b3\5$\23\2"+
		"\u00b2\u00b0\3\2\2\2\u00b2\u00b1\3\2\2\2\u00b3!\3\2\2\2\u00b4\u00b5\7"+
		"\61\2\2\u00b5\u00b9\5*\26\2\u00b6\u00b8\5\62\32\2\u00b7\u00b6\3\2\2\2"+
		"\u00b8\u00bb\3\2\2\2\u00b9\u00b7\3\2\2\2\u00b9\u00ba\3\2\2\2\u00ba\u00bc"+
		"\3\2\2\2\u00bb\u00b9\3\2\2\2\u00bc\u00bd\7\b\2\2\u00bd\u00be\7\3\2\2\u00be"+
		"\u00bf\7\61\2\2\u00bf\u00c0\7*\2\2\u00c0\u00c1\7\6\2\2\u00c1#\3\2\2\2"+
		"\u00c2\u00c6\7\61\2\2\u00c3\u00c5\5\62\32\2\u00c4\u00c3\3\2\2\2\u00c5"+
		"\u00c8\3\2\2\2\u00c6\u00c4\3\2\2\2\u00c6\u00c7\3\2\2\2\u00c7%\3\2\2\2"+
		"\u00c8\u00c6\3\2\2\2\u00c9\u00cc\5(\25\2\u00ca\u00cc\7\61\2\2\u00cb\u00c9"+
		"\3\2\2\2\u00cb\u00ca\3\2\2\2\u00cc\u00d1\3\2\2\2\u00cd\u00d0\5L\'\2\u00ce"+
		"\u00d0\5J&\2\u00cf\u00cd\3\2\2\2\u00cf\u00ce\3\2\2\2\u00d0\u00d3\3\2\2"+
		"\2\u00d1\u00cf\3\2\2\2\u00d1\u00d2\3\2\2\2\u00d2\'\3\2\2\2\u00d3\u00d1"+
		"\3\2\2\2\u00d4\u00d9\5:\36\2\u00d5\u00d9\5B\"\2\u00d6\u00d9\5<\37\2\u00d7"+
		"\u00d9\5> \2\u00d8\u00d4\3\2\2\2\u00d8\u00d5\3\2\2\2\u00d8\u00d6\3\2\2"+
		"\2\u00d8\u00d7\3\2\2\2\u00d9)\3\2\2\2\u00da\u00df\5\60\31\2\u00db\u00df"+
		"\5.\30\2\u00dc\u00df\5,\27\2\u00dd\u00df\7\61\2\2\u00de\u00da\3\2\2\2"+
		"\u00de\u00db\3\2\2\2\u00de\u00dc\3\2\2\2\u00de\u00dd\3\2\2\2\u00df+\3"+
		"\2\2\2\u00e0\u00e1\7\27\2\2\u00e1-\3\2\2\2\u00e2\u00e3\7\30\2\2\u00e3"+
		"/\3\2\2\2\u00e4\u00e5\7\31\2\2\u00e5\u00e6\7\32\2\2\u00e6\61\3\2\2\2\u00e7"+
		"\u00eb\5\64\33\2\u00e8\u00eb\5\66\34\2\u00e9\u00eb\58\35\2\u00ea\u00e7"+
		"\3\2\2\2\u00ea\u00e8\3\2\2\2\u00ea\u00e9\3\2\2\2\u00eb\63\3\2\2\2\u00ec"+
		"\u00f1\7\33\2\2\u00ed\u00f2\5&\24\2\u00ee\u00ef\7\34\2\2\u00ef\u00f0\7"+
		"\35\2\2\u00f0\u00f2\7\61\2\2\u00f1\u00ed\3\2\2\2\u00f1\u00ee\3\2\2\2\u00f2"+
		"\65\3\2\2\2\u00f3\u00f5\t\3\2\2\u00f4\u00f6\7\61\2\2\u00f5\u00f4\3\2\2"+
		"\2\u00f5\u00f6\3\2\2\2\u00f6\u00f7\3\2\2\2\u00f7\u00f8\7\3\2\2\u00f8\u00fd"+
		"\7\61\2\2\u00f9\u00fa\7\26\2\2\u00fa\u00fc\7\61\2\2\u00fb\u00f9\3\2\2"+
		"\2\u00fc\u00ff\3\2\2\2\u00fd\u00fb\3\2\2\2\u00fd\u00fe\3\2\2\2\u00fe\u0100"+
		"\3\2\2\2\u00ff\u00fd\3\2\2\2\u0100\u0101\7\6\2\2\u0101\67\3\2\2\2\u0102"+
		"\u0103\t\4\2\2\u0103\u0104\t\5\2\2\u01049\3\2\2\2\u0105\u0106\7\37\2\2"+
		"\u0106\u0107\7 \2\2\u0107;\3\2\2\2\u0108\u0109\7\31\2\2\u0109\u010a\7"+
		"\32\2\2\u010a=\3\2\2\2\u010b\u010c\7\34\2\2\u010c\u010d\7\3\2\2\u010d"+
		"\u0112\5@!\2\u010e\u010f\7\26\2\2\u010f\u0111\5@!\2\u0110\u010e\3\2\2"+
		"\2\u0111\u0114\3\2\2\2\u0112\u0110\3\2\2\2\u0112\u0113\3\2\2\2\u0113\u0115"+
		"\3\2\2\2\u0114\u0112\3\2\2\2\u0115\u0116\7\6\2\2\u0116?\3\2\2\2\u0117"+
		"\u0118\7\61\2\2\u0118\u0119\5&\24\2\u0119A\3\2\2\2\u011a\u0128\t\6\2\2"+
		"\u011b\u011c\7\3\2\2\u011c\u011d\5D#\2\u011d\u011e\7\6\2\2\u011e\u0129"+
		"\3\2\2\2\u011f\u0120\7\4\2\2\u0120\u0123\7*\2\2\u0121\u0122\7%\2\2\u0122"+
		"\u0124\7*\2\2\u0123\u0121\3\2\2\2\u0124\u0125\3\2\2\2\u0125\u0123\3\2"+
		"\2\2\u0125\u0126\3\2\2\2\u0126\u0127\3\2\2\2\u0127\u0129\7\5\2\2\u0128"+
		"\u011b\3\2\2\2\u0128\u011f\3\2\2\2\u0128\u0129\3\2\2\2\u0129C\3\2\2\2"+
		"\u012a\u012f\5F$\2\u012b\u012c\7\26\2\2\u012c\u012e\5F$\2\u012d\u012b"+
		"\3\2\2\2\u012e\u0131\3\2\2\2\u012f\u012d\3\2\2\2\u012f\u0130\3\2\2\2\u0130"+
		"E\3\2\2\2\u0131\u012f\3\2\2\2\u0132\u0133\7\61\2\2\u0133\u0134\7\4\2\2"+
		"\u0134\u0135\5H%\2\u0135\u0136\7\5\2\2\u0136G\3\2\2\2\u0137\u0139\7&\2"+
		"\2\u0138\u0137\3\2\2\2\u0138\u0139\3\2\2\2\u0139\u013a\3\2\2\2\u013a\u013b"+
		"\7*\2\2\u013bI\3\2\2\2\u013c\u013d\7\4\2\2\u013d\u013e\7\'\2\2\u013e\u013f"+
		"\5L\'\2\u013f\u0140\7\5\2\2\u0140K\3\2\2\2\u0141\u0144\7\4\2\2\u0142\u0145"+
		"\5N(\2\u0143\u0145\7*\2\2\u0144\u0142\3\2\2\2\u0144\u0143\3\2\2\2\u0145"+
		"\u0146\3\2\2\2\u0146\u0147\7\5\2\2\u0147M\3\2\2\2\u0148\u0149\5H%\2\u0149"+
		"\u014a\7(\2\2\u014a\u014b\5H%\2\u014bO\3\2\2\2\"X\\hlr{~\u0084\u0087\u008d"+
		"\u009b\u00a1\u00a7\u00ac\u00b2\u00b9\u00c6\u00cb\u00cf\u00d1\u00d8\u00de"+
		"\u00ea\u00f1\u00f5\u00fd\u0112\u0125\u0128\u012f\u0138\u0144";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}