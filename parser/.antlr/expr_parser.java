// Generated from c://Users//navar//GolandProjects//MiniGolang//parser//expr_parser.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class expr_parser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		PACKAGE=1, VAR=2, TYPE=3, FUNCTION=4, STRUCT=5, APPEND=6, LENGHT=7, CAP=8, 
		PRINT=9, PRINTLN=10, RETURN=11, BREAK=12, CONTINUE=13, IF=14, ELSE=15, 
		FOR=16, SWITCH=17, CASE=18, DEFAULT=19, ASSIGN=20, SHORT_DEC=21, SEMICOLON=22, 
		COLON=23, DOT=24, LPAREN=25, RPAREN=26, RBRACK=27, LBRACK=28, LBRACE=29, 
		RBRACE=30, COMMA=31, ADD=32, SUB=33, MULT=34, DIV=35, MOD=36, LSHIFT=37, 
		RSHIFT=38, AND=39, ANDNOT=40, PIPE=41, CARET=42, EQUALS=43, NOT_EQ=44, 
		LESS_THAN=45, LESS_THAN_OR_EQUALS=46, GREATER_THAN=47, GREATER_THAN_OR_EQUALS=48, 
		LOG_AND=49, LOG_OR=50, LOG_NOT=51, INCREMENT=52, DECREMENT=53, PLUS_ASSIGN=54, 
		AND_ASSIGN=55, MINUS_ASSIGN=56, OR_ASSIGN=57, TIMES_ASSIGN=58, XOR_ASSIGN=59, 
		LEFT_SHIFT_ASSIGN=60, RIGHT_SHIFT_ASSIGN=61, AND_NOT_ASSIGN=62, MOD_ASSIGN=63, 
		DIVIDE_ASSIGN=64, INT_LIT=65, IDENTIFIER=66, FLOAT_LIT=67, RUNE_LIT=68, 
		RAW_STRING_LIT=69, INTERPRETED_STRING_LIT=70, STRING_LIT=71, WS=72;
	public static final int
		RULE_root = 0, RULE_topDeclarationList = 1, RULE_variableDecl = 2, RULE_innerVarDecls = 3, 
		RULE_singleVarDecl = 4, RULE_singleVarDeclNoExps = 5, RULE_typeDecl = 6, 
		RULE_innerTypeDecls = 7, RULE_singleTypeDecl = 8, RULE_funcDecl = 9, RULE_funcFrontDecl = 10, 
		RULE_funcArgDecls = 11, RULE_declType = 12, RULE_sliceDeclType = 13, RULE_arrayDeclType = 14, 
		RULE_structDeclType = 15, RULE_structMemDecls = 16, RULE_identifierList = 17, 
		RULE_expression = 18, RULE_expressionList = 19, RULE_primaryExpression = 20, 
		RULE_operand = 21, RULE_literal = 22, RULE_index = 23, RULE_arguments = 24, 
		RULE_selector = 25, RULE_appendExpression = 26, RULE_lengthExpression = 27, 
		RULE_capExpression = 28, RULE_statementList = 29, RULE_block = 30, RULE_statement = 31, 
		RULE_simpleStatement = 32, RULE_assignmentStatement = 33, RULE_ifStatement = 34, 
		RULE_loop = 35, RULE_switch = 36, RULE_expressionCaseClauseList = 37, 
		RULE_expressionCaseClause = 38, RULE_expressionSwitchCase = 39;
	private static String[] makeRuleNames() {
		return new String[] {
			"root", "topDeclarationList", "variableDecl", "innerVarDecls", "singleVarDecl", 
			"singleVarDeclNoExps", "typeDecl", "innerTypeDecls", "singleTypeDecl", 
			"funcDecl", "funcFrontDecl", "funcArgDecls", "declType", "sliceDeclType", 
			"arrayDeclType", "structDeclType", "structMemDecls", "identifierList", 
			"expression", "expressionList", "primaryExpression", "operand", "literal", 
			"index", "arguments", "selector", "appendExpression", "lengthExpression", 
			"capExpression", "statementList", "block", "statement", "simpleStatement", 
			"assignmentStatement", "ifStatement", "loop", "switch", "expressionCaseClauseList", 
			"expressionCaseClause", "expressionSwitchCase"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'package'", "'var'", "'type'", "'func'", "'struct'", "'append'", 
			"'len'", "'cap'", "'print'", "'println'", "'return'", "'break'", "'continue'", 
			"'if'", "'else'", "'for'", "'switch'", "'case'", "'default'", "'='", 
			"':='", "';'", "':'", "'.'", "'('", "')'", "']'", "'['", "'{'", "'}'", 
			"','", "'+'", "'-'", "'*'", "'/'", "'%'", "'<<'", "'>>'", "'&'", "'&^'", 
			"'|'", "'^'", "'=='", "'!='", "'<'", "'<='", "'>'", "'>='", "'&&'", "'||'", 
			"'!'", "'++'", "'--'", "'+='", "'&='", "'-='", "'|='", "'*='", "'^='", 
			"'<<='", "'>>='", "'&^='", "'%='", "'/='"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "PACKAGE", "VAR", "TYPE", "FUNCTION", "STRUCT", "APPEND", "LENGHT", 
			"CAP", "PRINT", "PRINTLN", "RETURN", "BREAK", "CONTINUE", "IF", "ELSE", 
			"FOR", "SWITCH", "CASE", "DEFAULT", "ASSIGN", "SHORT_DEC", "SEMICOLON", 
			"COLON", "DOT", "LPAREN", "RPAREN", "RBRACK", "LBRACK", "LBRACE", "RBRACE", 
			"COMMA", "ADD", "SUB", "MULT", "DIV", "MOD", "LSHIFT", "RSHIFT", "AND", 
			"ANDNOT", "PIPE", "CARET", "EQUALS", "NOT_EQ", "LESS_THAN", "LESS_THAN_OR_EQUALS", 
			"GREATER_THAN", "GREATER_THAN_OR_EQUALS", "LOG_AND", "LOG_OR", "LOG_NOT", 
			"INCREMENT", "DECREMENT", "PLUS_ASSIGN", "AND_ASSIGN", "MINUS_ASSIGN", 
			"OR_ASSIGN", "TIMES_ASSIGN", "XOR_ASSIGN", "LEFT_SHIFT_ASSIGN", "RIGHT_SHIFT_ASSIGN", 
			"AND_NOT_ASSIGN", "MOD_ASSIGN", "DIVIDE_ASSIGN", "INT_LIT", "IDENTIFIER", 
			"FLOAT_LIT", "RUNE_LIT", "RAW_STRING_LIT", "INTERPRETED_STRING_LIT", 
			"STRING_LIT", "WS"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
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
	public String getGrammarFileName() { return "expr_parser.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public expr_parser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class RootContext extends ParserRuleContext {
		public TerminalNode PACKAGE() { return getToken(expr_parser.PACKAGE, 0); }
		public TerminalNode IDENTIFIER() { return getToken(expr_parser.IDENTIFIER, 0); }
		public TerminalNode SEMICOLON() { return getToken(expr_parser.SEMICOLON, 0); }
		public TopDeclarationListContext topDeclarationList() {
			return getRuleContext(TopDeclarationListContext.class,0);
		}
		public RootContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_root; }
	}

	public final RootContext root() throws RecognitionException {
		RootContext _localctx = new RootContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_root);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(80);
			match(PACKAGE);
			setState(81);
			match(IDENTIFIER);
			setState(82);
			match(SEMICOLON);
			setState(83);
			topDeclarationList();
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

	@SuppressWarnings("CheckReturnValue")
	public static class TopDeclarationListContext extends ParserRuleContext {
		public List<VariableDeclContext> variableDecl() {
			return getRuleContexts(VariableDeclContext.class);
		}
		public VariableDeclContext variableDecl(int i) {
			return getRuleContext(VariableDeclContext.class,i);
		}
		public List<TypeDeclContext> typeDecl() {
			return getRuleContexts(TypeDeclContext.class);
		}
		public TypeDeclContext typeDecl(int i) {
			return getRuleContext(TypeDeclContext.class,i);
		}
		public List<FuncDeclContext> funcDecl() {
			return getRuleContexts(FuncDeclContext.class);
		}
		public FuncDeclContext funcDecl(int i) {
			return getRuleContext(FuncDeclContext.class,i);
		}
		public TopDeclarationListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_topDeclarationList; }
	}

	public final TopDeclarationListContext topDeclarationList() throws RecognitionException {
		TopDeclarationListContext _localctx = new TopDeclarationListContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_topDeclarationList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(90);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 28L) != 0)) {
				{
				setState(88);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case VAR:
					{
					setState(85);
					variableDecl();
					}
					break;
				case TYPE:
					{
					setState(86);
					typeDecl();
					}
					break;
				case FUNCTION:
					{
					setState(87);
					funcDecl();
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				}
				setState(92);
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

	@SuppressWarnings("CheckReturnValue")
	public static class VariableDeclContext extends ParserRuleContext {
		public TerminalNode VAR() { return getToken(expr_parser.VAR, 0); }
		public TerminalNode SEMICOLON() { return getToken(expr_parser.SEMICOLON, 0); }
		public SingleVarDeclContext singleVarDecl() {
			return getRuleContext(SingleVarDeclContext.class,0);
		}
		public TerminalNode LPAREN() { return getToken(expr_parser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(expr_parser.RPAREN, 0); }
		public InnerVarDeclsContext innerVarDecls() {
			return getRuleContext(InnerVarDeclsContext.class,0);
		}
		public VariableDeclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_variableDecl; }
	}

	public final VariableDeclContext variableDecl() throws RecognitionException {
		VariableDeclContext _localctx = new VariableDeclContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_variableDecl);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(93);
			match(VAR);
			setState(100);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case IDENTIFIER:
				{
				setState(94);
				singleVarDecl();
				}
				break;
			case LPAREN:
				{
				setState(95);
				match(LPAREN);
				setState(97);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==IDENTIFIER) {
					{
					setState(96);
					innerVarDecls();
					}
				}

				setState(99);
				match(RPAREN);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(102);
			match(SEMICOLON);
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

	@SuppressWarnings("CheckReturnValue")
	public static class InnerVarDeclsContext extends ParserRuleContext {
		public List<SingleVarDeclContext> singleVarDecl() {
			return getRuleContexts(SingleVarDeclContext.class);
		}
		public SingleVarDeclContext singleVarDecl(int i) {
			return getRuleContext(SingleVarDeclContext.class,i);
		}
		public List<TerminalNode> SEMICOLON() { return getTokens(expr_parser.SEMICOLON); }
		public TerminalNode SEMICOLON(int i) {
			return getToken(expr_parser.SEMICOLON, i);
		}
		public InnerVarDeclsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_innerVarDecls; }
	}

	public final InnerVarDeclsContext innerVarDecls() throws RecognitionException {
		InnerVarDeclsContext _localctx = new InnerVarDeclsContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_innerVarDecls);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(107); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(104);
				singleVarDecl();
				setState(105);
				match(SEMICOLON);
				}
				}
				setState(109); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==IDENTIFIER );
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

	@SuppressWarnings("CheckReturnValue")
	public static class SingleVarDeclContext extends ParserRuleContext {
		public IdentifierListContext identifierList() {
			return getRuleContext(IdentifierListContext.class,0);
		}
		public DeclTypeContext declType() {
			return getRuleContext(DeclTypeContext.class,0);
		}
		public TerminalNode ASSIGN() { return getToken(expr_parser.ASSIGN, 0); }
		public ExpressionListContext expressionList() {
			return getRuleContext(ExpressionListContext.class,0);
		}
		public SingleVarDeclNoExpsContext singleVarDeclNoExps() {
			return getRuleContext(SingleVarDeclNoExpsContext.class,0);
		}
		public SingleVarDeclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_singleVarDecl; }
	}

	public final SingleVarDeclContext singleVarDecl() throws RecognitionException {
		SingleVarDeclContext _localctx = new SingleVarDeclContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_singleVarDecl);
		try {
			setState(121);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(111);
				identifierList();
				setState(112);
				declType();
				setState(113);
				match(ASSIGN);
				setState(114);
				expressionList();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(116);
				identifierList();
				setState(117);
				match(ASSIGN);
				setState(118);
				expressionList();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(120);
				singleVarDeclNoExps();
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

	@SuppressWarnings("CheckReturnValue")
	public static class SingleVarDeclNoExpsContext extends ParserRuleContext {
		public IdentifierListContext identifierList() {
			return getRuleContext(IdentifierListContext.class,0);
		}
		public DeclTypeContext declType() {
			return getRuleContext(DeclTypeContext.class,0);
		}
		public SingleVarDeclNoExpsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_singleVarDeclNoExps; }
	}

	public final SingleVarDeclNoExpsContext singleVarDeclNoExps() throws RecognitionException {
		SingleVarDeclNoExpsContext _localctx = new SingleVarDeclNoExpsContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_singleVarDeclNoExps);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(123);
			identifierList();
			setState(124);
			declType();
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

	@SuppressWarnings("CheckReturnValue")
	public static class TypeDeclContext extends ParserRuleContext {
		public TerminalNode TYPE() { return getToken(expr_parser.TYPE, 0); }
		public TerminalNode SEMICOLON() { return getToken(expr_parser.SEMICOLON, 0); }
		public SingleTypeDeclContext singleTypeDecl() {
			return getRuleContext(SingleTypeDeclContext.class,0);
		}
		public TerminalNode LPAREN() { return getToken(expr_parser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(expr_parser.RPAREN, 0); }
		public InnerTypeDeclsContext innerTypeDecls() {
			return getRuleContext(InnerTypeDeclsContext.class,0);
		}
		public TypeDeclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeDecl; }
	}

	public final TypeDeclContext typeDecl() throws RecognitionException {
		TypeDeclContext _localctx = new TypeDeclContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_typeDecl);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(126);
			match(TYPE);
			setState(133);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case IDENTIFIER:
				{
				setState(127);
				singleTypeDecl();
				}
				break;
			case LPAREN:
				{
				setState(128);
				match(LPAREN);
				setState(130);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==IDENTIFIER) {
					{
					setState(129);
					innerTypeDecls();
					}
				}

				setState(132);
				match(RPAREN);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(135);
			match(SEMICOLON);
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

	@SuppressWarnings("CheckReturnValue")
	public static class InnerTypeDeclsContext extends ParserRuleContext {
		public List<SingleTypeDeclContext> singleTypeDecl() {
			return getRuleContexts(SingleTypeDeclContext.class);
		}
		public SingleTypeDeclContext singleTypeDecl(int i) {
			return getRuleContext(SingleTypeDeclContext.class,i);
		}
		public List<TerminalNode> SEMICOLON() { return getTokens(expr_parser.SEMICOLON); }
		public TerminalNode SEMICOLON(int i) {
			return getToken(expr_parser.SEMICOLON, i);
		}
		public InnerTypeDeclsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_innerTypeDecls; }
	}

	public final InnerTypeDeclsContext innerTypeDecls() throws RecognitionException {
		InnerTypeDeclsContext _localctx = new InnerTypeDeclsContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_innerTypeDecls);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(137);
			singleTypeDecl();
			setState(138);
			match(SEMICOLON);
			setState(144);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==IDENTIFIER) {
				{
				{
				setState(139);
				singleTypeDecl();
				setState(140);
				match(SEMICOLON);
				}
				}
				setState(146);
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

	@SuppressWarnings("CheckReturnValue")
	public static class SingleTypeDeclContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(expr_parser.IDENTIFIER, 0); }
		public DeclTypeContext declType() {
			return getRuleContext(DeclTypeContext.class,0);
		}
		public SingleTypeDeclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_singleTypeDecl; }
	}

	public final SingleTypeDeclContext singleTypeDecl() throws RecognitionException {
		SingleTypeDeclContext _localctx = new SingleTypeDeclContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_singleTypeDecl);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(147);
			match(IDENTIFIER);
			setState(148);
			declType();
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

	@SuppressWarnings("CheckReturnValue")
	public static class FuncDeclContext extends ParserRuleContext {
		public FuncFrontDeclContext funcFrontDecl() {
			return getRuleContext(FuncFrontDeclContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(expr_parser.SEMICOLON, 0); }
		public FuncDeclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcDecl; }
	}

	public final FuncDeclContext funcDecl() throws RecognitionException {
		FuncDeclContext _localctx = new FuncDeclContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_funcDecl);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(150);
			funcFrontDecl();
			setState(151);
			block();
			setState(152);
			match(SEMICOLON);
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

	@SuppressWarnings("CheckReturnValue")
	public static class FuncFrontDeclContext extends ParserRuleContext {
		public TerminalNode FUNCTION() { return getToken(expr_parser.FUNCTION, 0); }
		public TerminalNode IDENTIFIER() { return getToken(expr_parser.IDENTIFIER, 0); }
		public TerminalNode LPAREN() { return getToken(expr_parser.LPAREN, 0); }
		public FuncArgDeclsContext funcArgDecls() {
			return getRuleContext(FuncArgDeclsContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(expr_parser.RPAREN, 0); }
		public DeclTypeContext declType() {
			return getRuleContext(DeclTypeContext.class,0);
		}
		public FuncFrontDeclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcFrontDecl; }
	}

	public final FuncFrontDeclContext funcFrontDecl() throws RecognitionException {
		FuncFrontDeclContext _localctx = new FuncFrontDeclContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_funcFrontDecl);
		try {
			setState(176);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(154);
				match(FUNCTION);
				setState(155);
				match(IDENTIFIER);
				setState(156);
				match(LPAREN);
				setState(157);
				funcArgDecls();
				setState(158);
				match(RPAREN);
				setState(159);
				declType();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(161);
				match(FUNCTION);
				setState(162);
				match(IDENTIFIER);
				setState(163);
				match(LPAREN);
				setState(164);
				funcArgDecls();
				setState(165);
				match(RPAREN);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(167);
				match(FUNCTION);
				setState(168);
				match(IDENTIFIER);
				setState(169);
				match(LPAREN);
				setState(170);
				match(RPAREN);
				setState(171);
				declType();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(172);
				match(FUNCTION);
				setState(173);
				match(IDENTIFIER);
				setState(174);
				match(LPAREN);
				setState(175);
				match(RPAREN);
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

	@SuppressWarnings("CheckReturnValue")
	public static class FuncArgDeclsContext extends ParserRuleContext {
		public List<SingleVarDeclNoExpsContext> singleVarDeclNoExps() {
			return getRuleContexts(SingleVarDeclNoExpsContext.class);
		}
		public SingleVarDeclNoExpsContext singleVarDeclNoExps(int i) {
			return getRuleContext(SingleVarDeclNoExpsContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(expr_parser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(expr_parser.COMMA, i);
		}
		public FuncArgDeclsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcArgDecls; }
	}

	public final FuncArgDeclsContext funcArgDecls() throws RecognitionException {
		FuncArgDeclsContext _localctx = new FuncArgDeclsContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_funcArgDecls);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(183);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,10,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(178);
					singleVarDeclNoExps();
					setState(179);
					match(COMMA);
					}
					} 
				}
				setState(185);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,10,_ctx);
			}
			setState(186);
			singleVarDeclNoExps();
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

	@SuppressWarnings("CheckReturnValue")
	public static class DeclTypeContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(expr_parser.LPAREN, 0); }
		public DeclTypeContext declType() {
			return getRuleContext(DeclTypeContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(expr_parser.RPAREN, 0); }
		public TerminalNode IDENTIFIER() { return getToken(expr_parser.IDENTIFIER, 0); }
		public SliceDeclTypeContext sliceDeclType() {
			return getRuleContext(SliceDeclTypeContext.class,0);
		}
		public ArrayDeclTypeContext arrayDeclType() {
			return getRuleContext(ArrayDeclTypeContext.class,0);
		}
		public StructDeclTypeContext structDeclType() {
			return getRuleContext(StructDeclTypeContext.class,0);
		}
		public DeclTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declType; }
	}

	public final DeclTypeContext declType() throws RecognitionException {
		DeclTypeContext _localctx = new DeclTypeContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_declType);
		try {
			setState(197);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(188);
				match(LPAREN);
				setState(189);
				declType();
				setState(190);
				match(RPAREN);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(192);
				match(IDENTIFIER);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(193);
				sliceDeclType();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(194);
				arrayDeclType();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(195);
				structDeclType();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
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

	@SuppressWarnings("CheckReturnValue")
	public static class SliceDeclTypeContext extends ParserRuleContext {
		public TerminalNode LBRACK() { return getToken(expr_parser.LBRACK, 0); }
		public TerminalNode RBRACK() { return getToken(expr_parser.RBRACK, 0); }
		public DeclTypeContext declType() {
			return getRuleContext(DeclTypeContext.class,0);
		}
		public SliceDeclTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sliceDeclType; }
	}

	public final SliceDeclTypeContext sliceDeclType() throws RecognitionException {
		SliceDeclTypeContext _localctx = new SliceDeclTypeContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_sliceDeclType);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(199);
			match(LBRACK);
			setState(200);
			match(RBRACK);
			setState(201);
			declType();
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

	@SuppressWarnings("CheckReturnValue")
	public static class ArrayDeclTypeContext extends ParserRuleContext {
		public TerminalNode LBRACK() { return getToken(expr_parser.LBRACK, 0); }
		public TerminalNode INT_LIT() { return getToken(expr_parser.INT_LIT, 0); }
		public TerminalNode RBRACK() { return getToken(expr_parser.RBRACK, 0); }
		public DeclTypeContext declType() {
			return getRuleContext(DeclTypeContext.class,0);
		}
		public ArrayDeclTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arrayDeclType; }
	}

	public final ArrayDeclTypeContext arrayDeclType() throws RecognitionException {
		ArrayDeclTypeContext _localctx = new ArrayDeclTypeContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_arrayDeclType);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(203);
			match(LBRACK);
			setState(204);
			match(INT_LIT);
			setState(205);
			match(RBRACK);
			setState(206);
			declType();
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

	@SuppressWarnings("CheckReturnValue")
	public static class StructDeclTypeContext extends ParserRuleContext {
		public TerminalNode STRUCT() { return getToken(expr_parser.STRUCT, 0); }
		public TerminalNode LBRACE() { return getToken(expr_parser.LBRACE, 0); }
		public TerminalNode RBRACE() { return getToken(expr_parser.RBRACE, 0); }
		public StructMemDeclsContext structMemDecls() {
			return getRuleContext(StructMemDeclsContext.class,0);
		}
		public StructDeclTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structDeclType; }
	}

	public final StructDeclTypeContext structDeclType() throws RecognitionException {
		StructDeclTypeContext _localctx = new StructDeclTypeContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_structDeclType);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(208);
			match(STRUCT);
			setState(209);
			match(LBRACE);
			setState(211);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==IDENTIFIER) {
				{
				setState(210);
				structMemDecls();
				}
			}

			setState(213);
			match(RBRACE);
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

	@SuppressWarnings("CheckReturnValue")
	public static class StructMemDeclsContext extends ParserRuleContext {
		public List<SingleVarDeclNoExpsContext> singleVarDeclNoExps() {
			return getRuleContexts(SingleVarDeclNoExpsContext.class);
		}
		public SingleVarDeclNoExpsContext singleVarDeclNoExps(int i) {
			return getRuleContext(SingleVarDeclNoExpsContext.class,i);
		}
		public List<TerminalNode> SEMICOLON() { return getTokens(expr_parser.SEMICOLON); }
		public TerminalNode SEMICOLON(int i) {
			return getToken(expr_parser.SEMICOLON, i);
		}
		public StructMemDeclsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structMemDecls; }
	}

	public final StructMemDeclsContext structMemDecls() throws RecognitionException {
		StructMemDeclsContext _localctx = new StructMemDeclsContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_structMemDecls);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(215);
			singleVarDeclNoExps();
			setState(216);
			match(SEMICOLON);
			setState(222);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==IDENTIFIER) {
				{
				{
				setState(217);
				singleVarDeclNoExps();
				setState(218);
				match(SEMICOLON);
				}
				}
				setState(224);
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

	@SuppressWarnings("CheckReturnValue")
	public static class IdentifierListContext extends ParserRuleContext {
		public List<TerminalNode> IDENTIFIER() { return getTokens(expr_parser.IDENTIFIER); }
		public TerminalNode IDENTIFIER(int i) {
			return getToken(expr_parser.IDENTIFIER, i);
		}
		public List<TerminalNode> COMMA() { return getTokens(expr_parser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(expr_parser.COMMA, i);
		}
		public IdentifierListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_identifierList; }
	}

	public final IdentifierListContext identifierList() throws RecognitionException {
		IdentifierListContext _localctx = new IdentifierListContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_identifierList);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(225);
			match(IDENTIFIER);
			setState(230);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,14,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(226);
					match(COMMA);
					setState(227);
					match(IDENTIFIER);
					}
					} 
				}
				setState(232);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,14,_ctx);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionContext extends ParserRuleContext {
		public PrimaryExpressionContext primaryExpression() {
			return getRuleContext(PrimaryExpressionContext.class,0);
		}
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode ADD() { return getToken(expr_parser.ADD, 0); }
		public TerminalNode SUB() { return getToken(expr_parser.SUB, 0); }
		public TerminalNode LOG_NOT() { return getToken(expr_parser.LOG_NOT, 0); }
		public TerminalNode CARET() { return getToken(expr_parser.CARET, 0); }
		public TerminalNode MULT() { return getToken(expr_parser.MULT, 0); }
		public TerminalNode DIV() { return getToken(expr_parser.DIV, 0); }
		public TerminalNode MOD() { return getToken(expr_parser.MOD, 0); }
		public TerminalNode LSHIFT() { return getToken(expr_parser.LSHIFT, 0); }
		public TerminalNode RSHIFT() { return getToken(expr_parser.RSHIFT, 0); }
		public TerminalNode AND() { return getToken(expr_parser.AND, 0); }
		public TerminalNode ANDNOT() { return getToken(expr_parser.ANDNOT, 0); }
		public TerminalNode PIPE() { return getToken(expr_parser.PIPE, 0); }
		public TerminalNode EQUALS() { return getToken(expr_parser.EQUALS, 0); }
		public TerminalNode NOT_EQ() { return getToken(expr_parser.NOT_EQ, 0); }
		public TerminalNode LESS_THAN() { return getToken(expr_parser.LESS_THAN, 0); }
		public TerminalNode LESS_THAN_OR_EQUALS() { return getToken(expr_parser.LESS_THAN_OR_EQUALS, 0); }
		public TerminalNode GREATER_THAN() { return getToken(expr_parser.GREATER_THAN, 0); }
		public TerminalNode GREATER_THAN_OR_EQUALS() { return getToken(expr_parser.GREATER_THAN_OR_EQUALS, 0); }
		public TerminalNode LOG_AND() { return getToken(expr_parser.LOG_AND, 0); }
		public TerminalNode LOG_OR() { return getToken(expr_parser.LOG_OR, 0); }
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	}

	public final ExpressionContext expression() throws RecognitionException {
		return expression(0);
	}

	private ExpressionContext expression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExpressionContext _localctx = new ExpressionContext(_ctx, _parentState);
		ExpressionContext _prevctx = _localctx;
		int _startState = 36;
		enterRecursionRule(_localctx, 36, RULE_expression, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(237);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case APPEND:
			case LENGHT:
			case CAP:
			case LPAREN:
			case INT_LIT:
			case IDENTIFIER:
			case FLOAT_LIT:
			case RAW_STRING_LIT:
			case INTERPRETED_STRING_LIT:
				{
				setState(234);
				primaryExpression(0);
				}
				break;
			case ADD:
			case SUB:
			case CARET:
			case LOG_NOT:
				{
				setState(235);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 2256210745098240L) != 0)) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(236);
				expression(1);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(244);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,16,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ExpressionContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_expression);
					setState(239);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(240);
					_la = _input.LA(1);
					if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 2251795518717952L) != 0)) ) {
					_errHandler.recoverInline(this);
					}
					else {
						if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
						_errHandler.reportMatch(this);
						consume();
					}
					setState(241);
					expression(3);
					}
					} 
				}
				setState(246);
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
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionListContext extends ParserRuleContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(expr_parser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(expr_parser.COMMA, i);
		}
		public ExpressionListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expressionList; }
	}

	public final ExpressionListContext expressionList() throws RecognitionException {
		ExpressionListContext _localctx = new ExpressionListContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_expressionList);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(247);
			expression(0);
			setState(252);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,17,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(248);
					match(COMMA);
					setState(249);
					expression(0);
					}
					} 
				}
				setState(254);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,17,_ctx);
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

	@SuppressWarnings("CheckReturnValue")
	public static class PrimaryExpressionContext extends ParserRuleContext {
		public OperandContext operand() {
			return getRuleContext(OperandContext.class,0);
		}
		public AppendExpressionContext appendExpression() {
			return getRuleContext(AppendExpressionContext.class,0);
		}
		public LengthExpressionContext lengthExpression() {
			return getRuleContext(LengthExpressionContext.class,0);
		}
		public CapExpressionContext capExpression() {
			return getRuleContext(CapExpressionContext.class,0);
		}
		public PrimaryExpressionContext primaryExpression() {
			return getRuleContext(PrimaryExpressionContext.class,0);
		}
		public SelectorContext selector() {
			return getRuleContext(SelectorContext.class,0);
		}
		public IndexContext index() {
			return getRuleContext(IndexContext.class,0);
		}
		public ArgumentsContext arguments() {
			return getRuleContext(ArgumentsContext.class,0);
		}
		public PrimaryExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_primaryExpression; }
	}

	public final PrimaryExpressionContext primaryExpression() throws RecognitionException {
		return primaryExpression(0);
	}

	private PrimaryExpressionContext primaryExpression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		PrimaryExpressionContext _localctx = new PrimaryExpressionContext(_ctx, _parentState);
		PrimaryExpressionContext _prevctx = _localctx;
		int _startState = 40;
		enterRecursionRule(_localctx, 40, RULE_primaryExpression, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(260);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LPAREN:
			case INT_LIT:
			case IDENTIFIER:
			case FLOAT_LIT:
			case RAW_STRING_LIT:
			case INTERPRETED_STRING_LIT:
				{
				setState(256);
				operand();
				}
				break;
			case APPEND:
				{
				setState(257);
				appendExpression();
				}
				break;
			case LENGHT:
				{
				setState(258);
				lengthExpression();
				}
				break;
			case CAP:
				{
				setState(259);
				capExpression();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(270);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,20,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(268);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,19,_ctx) ) {
					case 1:
						{
						_localctx = new PrimaryExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_primaryExpression);
						setState(262);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(263);
						selector();
						}
						break;
					case 2:
						{
						_localctx = new PrimaryExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_primaryExpression);
						setState(264);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(265);
						index();
						}
						break;
					case 3:
						{
						_localctx = new PrimaryExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_primaryExpression);
						setState(266);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(267);
						arguments();
						}
						break;
					}
					} 
				}
				setState(272);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,20,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class OperandContext extends ParserRuleContext {
		public LiteralContext literal() {
			return getRuleContext(LiteralContext.class,0);
		}
		public TerminalNode IDENTIFIER() { return getToken(expr_parser.IDENTIFIER, 0); }
		public TerminalNode LPAREN() { return getToken(expr_parser.LPAREN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(expr_parser.RPAREN, 0); }
		public OperandContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_operand; }
	}

	public final OperandContext operand() throws RecognitionException {
		OperandContext _localctx = new OperandContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_operand);
		try {
			setState(279);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT_LIT:
			case FLOAT_LIT:
			case RAW_STRING_LIT:
			case INTERPRETED_STRING_LIT:
				enterOuterAlt(_localctx, 1);
				{
				setState(273);
				literal();
				}
				break;
			case IDENTIFIER:
				enterOuterAlt(_localctx, 2);
				{
				setState(274);
				match(IDENTIFIER);
				}
				break;
			case LPAREN:
				enterOuterAlt(_localctx, 3);
				{
				setState(275);
				match(LPAREN);
				setState(276);
				expression(0);
				setState(277);
				match(RPAREN);
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

	@SuppressWarnings("CheckReturnValue")
	public static class LiteralContext extends ParserRuleContext {
		public TerminalNode INT_LIT() { return getToken(expr_parser.INT_LIT, 0); }
		public TerminalNode FLOAT_LIT() { return getToken(expr_parser.FLOAT_LIT, 0); }
		public TerminalNode RAW_STRING_LIT() { return getToken(expr_parser.RAW_STRING_LIT, 0); }
		public TerminalNode INTERPRETED_STRING_LIT() { return getToken(expr_parser.INTERPRETED_STRING_LIT, 0); }
		public LiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literal; }
	}

	public final LiteralContext literal() throws RecognitionException {
		LiteralContext _localctx = new LiteralContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_literal);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(281);
			_la = _input.LA(1);
			if ( !(((((_la - 65)) & ~0x3f) == 0 && ((1L << (_la - 65)) & 53L) != 0)) ) {
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

	@SuppressWarnings("CheckReturnValue")
	public static class IndexContext extends ParserRuleContext {
		public TerminalNode LBRACK() { return getToken(expr_parser.LBRACK, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode RBRACK() { return getToken(expr_parser.RBRACK, 0); }
		public IndexContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_index; }
	}

	public final IndexContext index() throws RecognitionException {
		IndexContext _localctx = new IndexContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_index);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(283);
			match(LBRACK);
			setState(284);
			expression(0);
			setState(285);
			match(RBRACK);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ArgumentsContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(expr_parser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(expr_parser.RPAREN, 0); }
		public ExpressionListContext expressionList() {
			return getRuleContext(ExpressionListContext.class,0);
		}
		public ArgumentsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arguments; }
	}

	public final ArgumentsContext arguments() throws RecognitionException {
		ArgumentsContext _localctx = new ArgumentsContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_arguments);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(287);
			match(LPAREN);
			setState(289);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 2256210778653120L) != 0) || ((((_la - 65)) & ~0x3f) == 0 && ((1L << (_la - 65)) & 55L) != 0)) {
				{
				setState(288);
				expressionList();
				}
			}

			setState(291);
			match(RPAREN);
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

	@SuppressWarnings("CheckReturnValue")
	public static class SelectorContext extends ParserRuleContext {
		public TerminalNode DOT() { return getToken(expr_parser.DOT, 0); }
		public TerminalNode IDENTIFIER() { return getToken(expr_parser.IDENTIFIER, 0); }
		public SelectorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_selector; }
	}

	public final SelectorContext selector() throws RecognitionException {
		SelectorContext _localctx = new SelectorContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_selector);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(293);
			match(DOT);
			setState(294);
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

	@SuppressWarnings("CheckReturnValue")
	public static class AppendExpressionContext extends ParserRuleContext {
		public TerminalNode APPEND() { return getToken(expr_parser.APPEND, 0); }
		public TerminalNode LPAREN() { return getToken(expr_parser.LPAREN, 0); }
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode COMMA() { return getToken(expr_parser.COMMA, 0); }
		public TerminalNode RPAREN() { return getToken(expr_parser.RPAREN, 0); }
		public AppendExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_appendExpression; }
	}

	public final AppendExpressionContext appendExpression() throws RecognitionException {
		AppendExpressionContext _localctx = new AppendExpressionContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_appendExpression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(296);
			match(APPEND);
			setState(297);
			match(LPAREN);
			setState(298);
			expression(0);
			setState(299);
			match(COMMA);
			setState(300);
			expression(0);
			setState(301);
			match(RPAREN);
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

	@SuppressWarnings("CheckReturnValue")
	public static class LengthExpressionContext extends ParserRuleContext {
		public TerminalNode LENGHT() { return getToken(expr_parser.LENGHT, 0); }
		public TerminalNode LPAREN() { return getToken(expr_parser.LPAREN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(expr_parser.RPAREN, 0); }
		public LengthExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_lengthExpression; }
	}

	public final LengthExpressionContext lengthExpression() throws RecognitionException {
		LengthExpressionContext _localctx = new LengthExpressionContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_lengthExpression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(303);
			match(LENGHT);
			setState(304);
			match(LPAREN);
			setState(305);
			expression(0);
			setState(306);
			match(RPAREN);
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

	@SuppressWarnings("CheckReturnValue")
	public static class CapExpressionContext extends ParserRuleContext {
		public TerminalNode CAP() { return getToken(expr_parser.CAP, 0); }
		public TerminalNode LPAREN() { return getToken(expr_parser.LPAREN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(expr_parser.RPAREN, 0); }
		public CapExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_capExpression; }
	}

	public final CapExpressionContext capExpression() throws RecognitionException {
		CapExpressionContext _localctx = new CapExpressionContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_capExpression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(308);
			match(CAP);
			setState(309);
			match(LPAREN);
			setState(310);
			expression(0);
			setState(311);
			match(RPAREN);
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

	@SuppressWarnings("CheckReturnValue")
	public static class StatementListContext extends ParserRuleContext {
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public StatementListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statementList; }
	}

	public final StatementListContext statementList() throws RecognitionException {
		StatementListContext _localctx = new StatementListContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_statementList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(316);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 2256211315752908L) != 0) || ((((_la - 65)) & ~0x3f) == 0 && ((1L << (_la - 65)) & 55L) != 0)) {
				{
				{
				setState(313);
				statement();
				}
				}
				setState(318);
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

	@SuppressWarnings("CheckReturnValue")
	public static class BlockContext extends ParserRuleContext {
		public TerminalNode LBRACE() { return getToken(expr_parser.LBRACE, 0); }
		public StatementListContext statementList() {
			return getRuleContext(StatementListContext.class,0);
		}
		public TerminalNode RBRACE() { return getToken(expr_parser.RBRACE, 0); }
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_block);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(319);
			match(LBRACE);
			setState(320);
			statementList();
			setState(321);
			match(RBRACE);
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

	@SuppressWarnings("CheckReturnValue")
	public static class StatementContext extends ParserRuleContext {
		public TerminalNode PRINT() { return getToken(expr_parser.PRINT, 0); }
		public TerminalNode LPAREN() { return getToken(expr_parser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(expr_parser.RPAREN, 0); }
		public TerminalNode COMMA() { return getToken(expr_parser.COMMA, 0); }
		public ExpressionListContext expressionList() {
			return getRuleContext(ExpressionListContext.class,0);
		}
		public TerminalNode PRINTLN() { return getToken(expr_parser.PRINTLN, 0); }
		public TerminalNode RETURN() { return getToken(expr_parser.RETURN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode BREAK() { return getToken(expr_parser.BREAK, 0); }
		public TerminalNode CONTINUE() { return getToken(expr_parser.CONTINUE, 0); }
		public SimpleStatementContext simpleStatement() {
			return getRuleContext(SimpleStatementContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public SwitchContext switch_() {
			return getRuleContext(SwitchContext.class,0);
		}
		public IfStatementContext ifStatement() {
			return getRuleContext(IfStatementContext.class,0);
		}
		public LoopContext loop() {
			return getRuleContext(LoopContext.class,0);
		}
		public TypeDeclContext typeDecl() {
			return getRuleContext(TypeDeclContext.class,0);
		}
		public VariableDeclContext variableDecl() {
			return getRuleContext(VariableDeclContext.class,0);
		}
		public StatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement; }
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_statement);
		int _la;
		try {
			setState(363);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case PRINT:
				enterOuterAlt(_localctx, 1);
				{
				setState(323);
				match(PRINT);
				setState(324);
				match(LPAREN);
				setState(326);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 2256210778653120L) != 0) || ((((_la - 65)) & ~0x3f) == 0 && ((1L << (_la - 65)) & 55L) != 0)) {
					{
					setState(325);
					expressionList();
					}
				}

				setState(328);
				match(RPAREN);
				setState(329);
				match(COMMA);
				}
				break;
			case PRINTLN:
				enterOuterAlt(_localctx, 2);
				{
				setState(330);
				match(PRINTLN);
				setState(331);
				match(LPAREN);
				setState(333);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 2256210778653120L) != 0) || ((((_la - 65)) & ~0x3f) == 0 && ((1L << (_la - 65)) & 55L) != 0)) {
					{
					setState(332);
					expressionList();
					}
				}

				setState(335);
				match(RPAREN);
				setState(336);
				match(COMMA);
				}
				break;
			case RETURN:
				enterOuterAlt(_localctx, 3);
				{
				setState(337);
				match(RETURN);
				setState(339);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 2256210778653120L) != 0) || ((((_la - 65)) & ~0x3f) == 0 && ((1L << (_la - 65)) & 55L) != 0)) {
					{
					setState(338);
					expression(0);
					}
				}

				setState(341);
				match(COMMA);
				}
				break;
			case BREAK:
				enterOuterAlt(_localctx, 4);
				{
				setState(342);
				match(BREAK);
				setState(343);
				match(COMMA);
				}
				break;
			case CONTINUE:
				enterOuterAlt(_localctx, 5);
				{
				setState(344);
				match(CONTINUE);
				setState(345);
				match(COMMA);
				}
				break;
			case APPEND:
			case LENGHT:
			case CAP:
			case LPAREN:
			case ADD:
			case SUB:
			case CARET:
			case LOG_NOT:
			case INT_LIT:
			case IDENTIFIER:
			case FLOAT_LIT:
			case RAW_STRING_LIT:
			case INTERPRETED_STRING_LIT:
				enterOuterAlt(_localctx, 6);
				{
				setState(346);
				simpleStatement();
				setState(347);
				match(COMMA);
				}
				break;
			case LBRACE:
				enterOuterAlt(_localctx, 7);
				{
				setState(349);
				block();
				setState(350);
				match(COMMA);
				}
				break;
			case SWITCH:
				enterOuterAlt(_localctx, 8);
				{
				setState(352);
				switch_();
				setState(353);
				match(COMMA);
				}
				break;
			case IF:
				enterOuterAlt(_localctx, 9);
				{
				setState(355);
				ifStatement();
				setState(356);
				match(COMMA);
				}
				break;
			case FOR:
				enterOuterAlt(_localctx, 10);
				{
				setState(358);
				loop();
				setState(359);
				match(COMMA);
				}
				break;
			case TYPE:
				enterOuterAlt(_localctx, 11);
				{
				setState(361);
				typeDecl();
				}
				break;
			case VAR:
				enterOuterAlt(_localctx, 12);
				{
				setState(362);
				variableDecl();
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

	@SuppressWarnings("CheckReturnValue")
	public static class SimpleStatementContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode INCREMENT() { return getToken(expr_parser.INCREMENT, 0); }
		public TerminalNode DECREMENT() { return getToken(expr_parser.DECREMENT, 0); }
		public AssignmentStatementContext assignmentStatement() {
			return getRuleContext(AssignmentStatementContext.class,0);
		}
		public List<ExpressionListContext> expressionList() {
			return getRuleContexts(ExpressionListContext.class);
		}
		public ExpressionListContext expressionList(int i) {
			return getRuleContext(ExpressionListContext.class,i);
		}
		public TerminalNode SHORT_DEC() { return getToken(expr_parser.SHORT_DEC, 0); }
		public SimpleStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_simpleStatement; }
	}

	public final SimpleStatementContext simpleStatement() throws RecognitionException {
		SimpleStatementContext _localctx = new SimpleStatementContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_simpleStatement);
		int _la;
		try {
			setState(374);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,29,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(365);
				expression(0);
				setState(367);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==INCREMENT || _la==DECREMENT) {
					{
					setState(366);
					_la = _input.LA(1);
					if ( !(_la==INCREMENT || _la==DECREMENT) ) {
					_errHandler.recoverInline(this);
					}
					else {
						if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
						_errHandler.reportMatch(this);
						consume();
					}
					}
				}

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(369);
				assignmentStatement();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(370);
				expressionList();
				setState(371);
				match(SHORT_DEC);
				setState(372);
				expressionList();
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

	@SuppressWarnings("CheckReturnValue")
	public static class AssignmentStatementContext extends ParserRuleContext {
		public List<ExpressionListContext> expressionList() {
			return getRuleContexts(ExpressionListContext.class);
		}
		public ExpressionListContext expressionList(int i) {
			return getRuleContext(ExpressionListContext.class,i);
		}
		public TerminalNode ASSIGN() { return getToken(expr_parser.ASSIGN, 0); }
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode PLUS_ASSIGN() { return getToken(expr_parser.PLUS_ASSIGN, 0); }
		public TerminalNode AND_ASSIGN() { return getToken(expr_parser.AND_ASSIGN, 0); }
		public TerminalNode MINUS_ASSIGN() { return getToken(expr_parser.MINUS_ASSIGN, 0); }
		public TerminalNode OR_ASSIGN() { return getToken(expr_parser.OR_ASSIGN, 0); }
		public TerminalNode TIMES_ASSIGN() { return getToken(expr_parser.TIMES_ASSIGN, 0); }
		public TerminalNode XOR_ASSIGN() { return getToken(expr_parser.XOR_ASSIGN, 0); }
		public TerminalNode LEFT_SHIFT_ASSIGN() { return getToken(expr_parser.LEFT_SHIFT_ASSIGN, 0); }
		public TerminalNode RIGHT_SHIFT_ASSIGN() { return getToken(expr_parser.RIGHT_SHIFT_ASSIGN, 0); }
		public TerminalNode AND_NOT_ASSIGN() { return getToken(expr_parser.AND_NOT_ASSIGN, 0); }
		public TerminalNode MOD_ASSIGN() { return getToken(expr_parser.MOD_ASSIGN, 0); }
		public TerminalNode DIVIDE_ASSIGN() { return getToken(expr_parser.DIVIDE_ASSIGN, 0); }
		public AssignmentStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignmentStatement; }
	}

	public final AssignmentStatementContext assignmentStatement() throws RecognitionException {
		AssignmentStatementContext _localctx = new AssignmentStatementContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_assignmentStatement);
		int _la;
		try {
			setState(384);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,30,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(376);
				expressionList();
				setState(377);
				match(ASSIGN);
				setState(378);
				expressionList();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(380);
				expression(0);
				setState(381);
				_la = _input.LA(1);
				if ( !(((((_la - 54)) & ~0x3f) == 0 && ((1L << (_la - 54)) & 2047L) != 0)) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(382);
				expression(0);
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

	@SuppressWarnings("CheckReturnValue")
	public static class IfStatementContext extends ParserRuleContext {
		public TerminalNode IF() { return getToken(expr_parser.IF, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public List<BlockContext> block() {
			return getRuleContexts(BlockContext.class);
		}
		public BlockContext block(int i) {
			return getRuleContext(BlockContext.class,i);
		}
		public TerminalNode ELSE() { return getToken(expr_parser.ELSE, 0); }
		public IfStatementContext ifStatement() {
			return getRuleContext(IfStatementContext.class,0);
		}
		public SimpleStatementContext simpleStatement() {
			return getRuleContext(SimpleStatementContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(expr_parser.SEMICOLON, 0); }
		public IfStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifStatement; }
	}

	public final IfStatementContext ifStatement() throws RecognitionException {
		IfStatementContext _localctx = new IfStatementContext(_ctx, getState());
		enterRule(_localctx, 68, RULE_ifStatement);
		try {
			setState(424);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,31,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(386);
				match(IF);
				setState(387);
				expression(0);
				setState(388);
				block();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(390);
				match(IF);
				setState(391);
				expression(0);
				setState(392);
				block();
				setState(393);
				match(ELSE);
				setState(394);
				ifStatement();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(396);
				match(IF);
				setState(397);
				expression(0);
				setState(398);
				block();
				setState(399);
				match(ELSE);
				setState(400);
				block();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(402);
				match(IF);
				setState(403);
				simpleStatement();
				setState(404);
				match(SEMICOLON);
				setState(405);
				expression(0);
				setState(406);
				block();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(408);
				match(IF);
				setState(409);
				simpleStatement();
				setState(410);
				match(SEMICOLON);
				setState(411);
				expression(0);
				setState(412);
				block();
				setState(413);
				match(ELSE);
				setState(414);
				ifStatement();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(416);
				match(IF);
				setState(417);
				simpleStatement();
				setState(418);
				match(SEMICOLON);
				setState(419);
				expression(0);
				setState(420);
				block();
				setState(421);
				match(ELSE);
				setState(422);
				block();
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

	@SuppressWarnings("CheckReturnValue")
	public static class LoopContext extends ParserRuleContext {
		public TerminalNode FOR() { return getToken(expr_parser.FOR, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public List<SimpleStatementContext> simpleStatement() {
			return getRuleContexts(SimpleStatementContext.class);
		}
		public SimpleStatementContext simpleStatement(int i) {
			return getRuleContext(SimpleStatementContext.class,i);
		}
		public List<TerminalNode> SEMICOLON() { return getTokens(expr_parser.SEMICOLON); }
		public TerminalNode SEMICOLON(int i) {
			return getToken(expr_parser.SEMICOLON, i);
		}
		public LoopContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_loop; }
	}

	public final LoopContext loop() throws RecognitionException {
		LoopContext _localctx = new LoopContext(_ctx, getState());
		enterRule(_localctx, 70, RULE_loop);
		try {
			setState(447);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,32,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(426);
				match(FOR);
				setState(427);
				block();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(428);
				match(FOR);
				setState(429);
				expression(0);
				setState(430);
				block();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(432);
				match(FOR);
				setState(433);
				simpleStatement();
				setState(434);
				match(SEMICOLON);
				setState(435);
				expression(0);
				setState(436);
				match(SEMICOLON);
				setState(437);
				simpleStatement();
				setState(438);
				block();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(440);
				match(FOR);
				setState(441);
				simpleStatement();
				setState(442);
				match(SEMICOLON);
				setState(443);
				match(SEMICOLON);
				setState(444);
				simpleStatement();
				setState(445);
				block();
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

	@SuppressWarnings("CheckReturnValue")
	public static class SwitchContext extends ParserRuleContext {
		public TerminalNode SWITCH() { return getToken(expr_parser.SWITCH, 0); }
		public TerminalNode LBRACE() { return getToken(expr_parser.LBRACE, 0); }
		public ExpressionCaseClauseListContext expressionCaseClauseList() {
			return getRuleContext(ExpressionCaseClauseListContext.class,0);
		}
		public TerminalNode RBRACE() { return getToken(expr_parser.RBRACE, 0); }
		public SimpleStatementContext simpleStatement() {
			return getRuleContext(SimpleStatementContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(expr_parser.SEMICOLON, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public SwitchContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switch; }
	}

	public final SwitchContext switch_() throws RecognitionException {
		SwitchContext _localctx = new SwitchContext(_ctx, getState());
		enterRule(_localctx, 72, RULE_switch);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(449);
			match(SWITCH);
			setState(451);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,33,_ctx) ) {
			case 1:
				{
				setState(450);
				simpleStatement();
				}
				break;
			}
			setState(454);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==SEMICOLON) {
				{
				setState(453);
				match(SEMICOLON);
				}
			}

			setState(457);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 2256210778653120L) != 0) || ((((_la - 65)) & ~0x3f) == 0 && ((1L << (_la - 65)) & 55L) != 0)) {
				{
				setState(456);
				expression(0);
				}
			}

			setState(459);
			match(LBRACE);
			setState(460);
			expressionCaseClauseList();
			setState(461);
			match(RBRACE);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionCaseClauseListContext extends ParserRuleContext {
		public List<ExpressionCaseClauseContext> expressionCaseClause() {
			return getRuleContexts(ExpressionCaseClauseContext.class);
		}
		public ExpressionCaseClauseContext expressionCaseClause(int i) {
			return getRuleContext(ExpressionCaseClauseContext.class,i);
		}
		public ExpressionCaseClauseListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expressionCaseClauseList; }
	}

	public final ExpressionCaseClauseListContext expressionCaseClauseList() throws RecognitionException {
		ExpressionCaseClauseListContext _localctx = new ExpressionCaseClauseListContext(_ctx, getState());
		enterRule(_localctx, 74, RULE_expressionCaseClauseList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(466);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==CASE || _la==DEFAULT) {
				{
				{
				setState(463);
				expressionCaseClause();
				}
				}
				setState(468);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionCaseClauseContext extends ParserRuleContext {
		public ExpressionSwitchCaseContext expressionSwitchCase() {
			return getRuleContext(ExpressionSwitchCaseContext.class,0);
		}
		public TerminalNode COLON() { return getToken(expr_parser.COLON, 0); }
		public StatementListContext statementList() {
			return getRuleContext(StatementListContext.class,0);
		}
		public ExpressionCaseClauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expressionCaseClause; }
	}

	public final ExpressionCaseClauseContext expressionCaseClause() throws RecognitionException {
		ExpressionCaseClauseContext _localctx = new ExpressionCaseClauseContext(_ctx, getState());
		enterRule(_localctx, 76, RULE_expressionCaseClause);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(469);
			expressionSwitchCase();
			setState(470);
			match(COLON);
			setState(471);
			statementList();
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

	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionSwitchCaseContext extends ParserRuleContext {
		public TerminalNode CASE() { return getToken(expr_parser.CASE, 0); }
		public ExpressionListContext expressionList() {
			return getRuleContext(ExpressionListContext.class,0);
		}
		public TerminalNode DEFAULT() { return getToken(expr_parser.DEFAULT, 0); }
		public ExpressionSwitchCaseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expressionSwitchCase; }
	}

	public final ExpressionSwitchCaseContext expressionSwitchCase() throws RecognitionException {
		ExpressionSwitchCaseContext _localctx = new ExpressionSwitchCaseContext(_ctx, getState());
		enterRule(_localctx, 78, RULE_expressionSwitchCase);
		try {
			setState(476);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case CASE:
				enterOuterAlt(_localctx, 1);
				{
				setState(473);
				match(CASE);
				setState(474);
				expressionList();
				}
				break;
			case DEFAULT:
				enterOuterAlt(_localctx, 2);
				{
				setState(475);
				match(DEFAULT);
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

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 18:
			return expression_sempred((ExpressionContext)_localctx, predIndex);
		case 20:
			return primaryExpression_sempred((PrimaryExpressionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expression_sempred(ExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean primaryExpression_sempred(PrimaryExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 1:
			return precpred(_ctx, 6);
		case 2:
			return precpred(_ctx, 5);
		case 3:
			return precpred(_ctx, 4);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001H\u01df\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002\u0018\u0007\u0018"+
		"\u0002\u0019\u0007\u0019\u0002\u001a\u0007\u001a\u0002\u001b\u0007\u001b"+
		"\u0002\u001c\u0007\u001c\u0002\u001d\u0007\u001d\u0002\u001e\u0007\u001e"+
		"\u0002\u001f\u0007\u001f\u0002 \u0007 \u0002!\u0007!\u0002\"\u0007\"\u0002"+
		"#\u0007#\u0002$\u0007$\u0002%\u0007%\u0002&\u0007&\u0002\'\u0007\'\u0001"+
		"\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0005\u0001Y\b\u0001\n\u0001\f\u0001\\\t\u0001\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0003\u0002b\b\u0002\u0001"+
		"\u0002\u0003\u0002e\b\u0002\u0001\u0002\u0001\u0002\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0004\u0003l\b\u0003\u000b\u0003\f\u0003m\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0003\u0004z\b\u0004\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0003\u0006\u0083\b\u0006\u0001\u0006\u0003\u0006\u0086\b\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001"+
		"\u0007\u0005\u0007\u008f\b\u0007\n\u0007\f\u0007\u0092\t\u0007\u0001\b"+
		"\u0001\b\u0001\b\u0001\t\u0001\t\u0001\t\u0001\t\u0001\n\u0001\n\u0001"+
		"\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001"+
		"\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001"+
		"\n\u0001\n\u0003\n\u00b1\b\n\u0001\u000b\u0001\u000b\u0001\u000b\u0005"+
		"\u000b\u00b6\b\u000b\n\u000b\f\u000b\u00b9\t\u000b\u0001\u000b\u0001\u000b"+
		"\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001"+
		"\f\u0003\f\u00c6\b\f\u0001\r\u0001\r\u0001\r\u0001\r\u0001\u000e\u0001"+
		"\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000f\u0001\u000f\u0001"+
		"\u000f\u0003\u000f\u00d4\b\u000f\u0001\u000f\u0001\u000f\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0005\u0010\u00dd\b\u0010\n"+
		"\u0010\f\u0010\u00e0\t\u0010\u0001\u0011\u0001\u0011\u0001\u0011\u0005"+
		"\u0011\u00e5\b\u0011\n\u0011\f\u0011\u00e8\t\u0011\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0003\u0012\u00ee\b\u0012\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0005\u0012\u00f3\b\u0012\n\u0012\f\u0012\u00f6\t\u0012\u0001"+
		"\u0013\u0001\u0013\u0001\u0013\u0005\u0013\u00fb\b\u0013\n\u0013\f\u0013"+
		"\u00fe\t\u0013\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014"+
		"\u0003\u0014\u0105\b\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014"+
		"\u0001\u0014\u0001\u0014\u0005\u0014\u010d\b\u0014\n\u0014\f\u0014\u0110"+
		"\t\u0014\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001"+
		"\u0015\u0003\u0015\u0118\b\u0015\u0001\u0016\u0001\u0016\u0001\u0017\u0001"+
		"\u0017\u0001\u0017\u0001\u0017\u0001\u0018\u0001\u0018\u0003\u0018\u0122"+
		"\b\u0018\u0001\u0018\u0001\u0018\u0001\u0019\u0001\u0019\u0001\u0019\u0001"+
		"\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0001"+
		"\u001a\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0001"+
		"\u001c\u0001\u001c\u0001\u001c\u0001\u001c\u0001\u001c\u0001\u001d\u0005"+
		"\u001d\u013b\b\u001d\n\u001d\f\u001d\u013e\t\u001d\u0001\u001e\u0001\u001e"+
		"\u0001\u001e\u0001\u001e\u0001\u001f\u0001\u001f\u0001\u001f\u0003\u001f"+
		"\u0147\b\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f"+
		"\u0003\u001f\u014e\b\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f"+
		"\u0003\u001f\u0154\b\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f"+
		"\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f"+
		"\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f"+
		"\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f"+
		"\u0003\u001f\u016c\b\u001f\u0001 \u0001 \u0003 \u0170\b \u0001 \u0001"+
		" \u0001 \u0001 \u0001 \u0003 \u0177\b \u0001!\u0001!\u0001!\u0001!\u0001"+
		"!\u0001!\u0001!\u0001!\u0003!\u0181\b!\u0001\"\u0001\"\u0001\"\u0001\""+
		"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001"+
		"\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001"+
		"\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001"+
		"\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0003\"\u01a9"+
		"\b\"\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001"+
		"#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001"+
		"#\u0001#\u0003#\u01c0\b#\u0001$\u0001$\u0003$\u01c4\b$\u0001$\u0003$\u01c7"+
		"\b$\u0001$\u0003$\u01ca\b$\u0001$\u0001$\u0001$\u0001$\u0001%\u0005%\u01d1"+
		"\b%\n%\f%\u01d4\t%\u0001&\u0001&\u0001&\u0001&\u0001\'\u0001\'\u0001\'"+
		"\u0003\'\u01dd\b\'\u0001\'\u0000\u0002$((\u0000\u0002\u0004\u0006\b\n"+
		"\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c\u001e \"$&(*,.0246"+
		"8:<>@BDFHJLN\u0000\u0005\u0003\u0000 !**33\u0001\u0000 2\u0003\u0000A"+
		"ACCEF\u0001\u000045\u0001\u00006@\u01f9\u0000P\u0001\u0000\u0000\u0000"+
		"\u0002Z\u0001\u0000\u0000\u0000\u0004]\u0001\u0000\u0000\u0000\u0006k"+
		"\u0001\u0000\u0000\u0000\by\u0001\u0000\u0000\u0000\n{\u0001\u0000\u0000"+
		"\u0000\f~\u0001\u0000\u0000\u0000\u000e\u0089\u0001\u0000\u0000\u0000"+
		"\u0010\u0093\u0001\u0000\u0000\u0000\u0012\u0096\u0001\u0000\u0000\u0000"+
		"\u0014\u00b0\u0001\u0000\u0000\u0000\u0016\u00b7\u0001\u0000\u0000\u0000"+
		"\u0018\u00c5\u0001\u0000\u0000\u0000\u001a\u00c7\u0001\u0000\u0000\u0000"+
		"\u001c\u00cb\u0001\u0000\u0000\u0000\u001e\u00d0\u0001\u0000\u0000\u0000"+
		" \u00d7\u0001\u0000\u0000\u0000\"\u00e1\u0001\u0000\u0000\u0000$\u00ed"+
		"\u0001\u0000\u0000\u0000&\u00f7\u0001\u0000\u0000\u0000(\u0104\u0001\u0000"+
		"\u0000\u0000*\u0117\u0001\u0000\u0000\u0000,\u0119\u0001\u0000\u0000\u0000"+
		".\u011b\u0001\u0000\u0000\u00000\u011f\u0001\u0000\u0000\u00002\u0125"+
		"\u0001\u0000\u0000\u00004\u0128\u0001\u0000\u0000\u00006\u012f\u0001\u0000"+
		"\u0000\u00008\u0134\u0001\u0000\u0000\u0000:\u013c\u0001\u0000\u0000\u0000"+
		"<\u013f\u0001\u0000\u0000\u0000>\u016b\u0001\u0000\u0000\u0000@\u0176"+
		"\u0001\u0000\u0000\u0000B\u0180\u0001\u0000\u0000\u0000D\u01a8\u0001\u0000"+
		"\u0000\u0000F\u01bf\u0001\u0000\u0000\u0000H\u01c1\u0001\u0000\u0000\u0000"+
		"J\u01d2\u0001\u0000\u0000\u0000L\u01d5\u0001\u0000\u0000\u0000N\u01dc"+
		"\u0001\u0000\u0000\u0000PQ\u0005\u0001\u0000\u0000QR\u0005B\u0000\u0000"+
		"RS\u0005\u0016\u0000\u0000ST\u0003\u0002\u0001\u0000T\u0001\u0001\u0000"+
		"\u0000\u0000UY\u0003\u0004\u0002\u0000VY\u0003\f\u0006\u0000WY\u0003\u0012"+
		"\t\u0000XU\u0001\u0000\u0000\u0000XV\u0001\u0000\u0000\u0000XW\u0001\u0000"+
		"\u0000\u0000Y\\\u0001\u0000\u0000\u0000ZX\u0001\u0000\u0000\u0000Z[\u0001"+
		"\u0000\u0000\u0000[\u0003\u0001\u0000\u0000\u0000\\Z\u0001\u0000\u0000"+
		"\u0000]d\u0005\u0002\u0000\u0000^e\u0003\b\u0004\u0000_a\u0005\u0019\u0000"+
		"\u0000`b\u0003\u0006\u0003\u0000a`\u0001\u0000\u0000\u0000ab\u0001\u0000"+
		"\u0000\u0000bc\u0001\u0000\u0000\u0000ce\u0005\u001a\u0000\u0000d^\u0001"+
		"\u0000\u0000\u0000d_\u0001\u0000\u0000\u0000ef\u0001\u0000\u0000\u0000"+
		"fg\u0005\u0016\u0000\u0000g\u0005\u0001\u0000\u0000\u0000hi\u0003\b\u0004"+
		"\u0000ij\u0005\u0016\u0000\u0000jl\u0001\u0000\u0000\u0000kh\u0001\u0000"+
		"\u0000\u0000lm\u0001\u0000\u0000\u0000mk\u0001\u0000\u0000\u0000mn\u0001"+
		"\u0000\u0000\u0000n\u0007\u0001\u0000\u0000\u0000op\u0003\"\u0011\u0000"+
		"pq\u0003\u0018\f\u0000qr\u0005\u0014\u0000\u0000rs\u0003&\u0013\u0000"+
		"sz\u0001\u0000\u0000\u0000tu\u0003\"\u0011\u0000uv\u0005\u0014\u0000\u0000"+
		"vw\u0003&\u0013\u0000wz\u0001\u0000\u0000\u0000xz\u0003\n\u0005\u0000"+
		"yo\u0001\u0000\u0000\u0000yt\u0001\u0000\u0000\u0000yx\u0001\u0000\u0000"+
		"\u0000z\t\u0001\u0000\u0000\u0000{|\u0003\"\u0011\u0000|}\u0003\u0018"+
		"\f\u0000}\u000b\u0001\u0000\u0000\u0000~\u0085\u0005\u0003\u0000\u0000"+
		"\u007f\u0086\u0003\u0010\b\u0000\u0080\u0082\u0005\u0019\u0000\u0000\u0081"+
		"\u0083\u0003\u000e\u0007\u0000\u0082\u0081\u0001\u0000\u0000\u0000\u0082"+
		"\u0083\u0001\u0000\u0000\u0000\u0083\u0084\u0001\u0000\u0000\u0000\u0084"+
		"\u0086\u0005\u001a\u0000\u0000\u0085\u007f\u0001\u0000\u0000\u0000\u0085"+
		"\u0080\u0001\u0000\u0000\u0000\u0086\u0087\u0001\u0000\u0000\u0000\u0087"+
		"\u0088\u0005\u0016\u0000\u0000\u0088\r\u0001\u0000\u0000\u0000\u0089\u008a"+
		"\u0003\u0010\b\u0000\u008a\u0090\u0005\u0016\u0000\u0000\u008b\u008c\u0003"+
		"\u0010\b\u0000\u008c\u008d\u0005\u0016\u0000\u0000\u008d\u008f\u0001\u0000"+
		"\u0000\u0000\u008e\u008b\u0001\u0000\u0000\u0000\u008f\u0092\u0001\u0000"+
		"\u0000\u0000\u0090\u008e\u0001\u0000\u0000\u0000\u0090\u0091\u0001\u0000"+
		"\u0000\u0000\u0091\u000f\u0001\u0000\u0000\u0000\u0092\u0090\u0001\u0000"+
		"\u0000\u0000\u0093\u0094\u0005B\u0000\u0000\u0094\u0095\u0003\u0018\f"+
		"\u0000\u0095\u0011\u0001\u0000\u0000\u0000\u0096\u0097\u0003\u0014\n\u0000"+
		"\u0097\u0098\u0003<\u001e\u0000\u0098\u0099\u0005\u0016\u0000\u0000\u0099"+
		"\u0013\u0001\u0000\u0000\u0000\u009a\u009b\u0005\u0004\u0000\u0000\u009b"+
		"\u009c\u0005B\u0000\u0000\u009c\u009d\u0005\u0019\u0000\u0000\u009d\u009e"+
		"\u0003\u0016\u000b\u0000\u009e\u009f\u0005\u001a\u0000\u0000\u009f\u00a0"+
		"\u0003\u0018\f\u0000\u00a0\u00b1\u0001\u0000\u0000\u0000\u00a1\u00a2\u0005"+
		"\u0004\u0000\u0000\u00a2\u00a3\u0005B\u0000\u0000\u00a3\u00a4\u0005\u0019"+
		"\u0000\u0000\u00a4\u00a5\u0003\u0016\u000b\u0000\u00a5\u00a6\u0005\u001a"+
		"\u0000\u0000\u00a6\u00b1\u0001\u0000\u0000\u0000\u00a7\u00a8\u0005\u0004"+
		"\u0000\u0000\u00a8\u00a9\u0005B\u0000\u0000\u00a9\u00aa\u0005\u0019\u0000"+
		"\u0000\u00aa\u00ab\u0005\u001a\u0000\u0000\u00ab\u00b1\u0003\u0018\f\u0000"+
		"\u00ac\u00ad\u0005\u0004\u0000\u0000\u00ad\u00ae\u0005B\u0000\u0000\u00ae"+
		"\u00af\u0005\u0019\u0000\u0000\u00af\u00b1\u0005\u001a\u0000\u0000\u00b0"+
		"\u009a\u0001\u0000\u0000\u0000\u00b0\u00a1\u0001\u0000\u0000\u0000\u00b0"+
		"\u00a7\u0001\u0000\u0000\u0000\u00b0\u00ac\u0001\u0000\u0000\u0000\u00b1"+
		"\u0015\u0001\u0000\u0000\u0000\u00b2\u00b3\u0003\n\u0005\u0000\u00b3\u00b4"+
		"\u0005\u001f\u0000\u0000\u00b4\u00b6\u0001\u0000\u0000\u0000\u00b5\u00b2"+
		"\u0001\u0000\u0000\u0000\u00b6\u00b9\u0001\u0000\u0000\u0000\u00b7\u00b5"+
		"\u0001\u0000\u0000\u0000\u00b7\u00b8\u0001\u0000\u0000\u0000\u00b8\u00ba"+
		"\u0001\u0000\u0000\u0000\u00b9\u00b7\u0001\u0000\u0000\u0000\u00ba\u00bb"+
		"\u0003\n\u0005\u0000\u00bb\u0017\u0001\u0000\u0000\u0000\u00bc\u00bd\u0005"+
		"\u0019\u0000\u0000\u00bd\u00be\u0003\u0018\f\u0000\u00be\u00bf\u0005\u001a"+
		"\u0000\u0000\u00bf\u00c6\u0001\u0000\u0000\u0000\u00c0\u00c6\u0005B\u0000"+
		"\u0000\u00c1\u00c6\u0003\u001a\r\u0000\u00c2\u00c6\u0003\u001c\u000e\u0000"+
		"\u00c3\u00c6\u0003\u001e\u000f\u0000\u00c4\u00c6\u0001\u0000\u0000\u0000"+
		"\u00c5\u00bc\u0001\u0000\u0000\u0000\u00c5\u00c0\u0001\u0000\u0000\u0000"+
		"\u00c5\u00c1\u0001\u0000\u0000\u0000\u00c5\u00c2\u0001\u0000\u0000\u0000"+
		"\u00c5\u00c3\u0001\u0000\u0000\u0000\u00c5\u00c4\u0001\u0000\u0000\u0000"+
		"\u00c6\u0019\u0001\u0000\u0000\u0000\u00c7\u00c8\u0005\u001c\u0000\u0000"+
		"\u00c8\u00c9\u0005\u001b\u0000\u0000\u00c9\u00ca\u0003\u0018\f\u0000\u00ca"+
		"\u001b\u0001\u0000\u0000\u0000\u00cb\u00cc\u0005\u001c\u0000\u0000\u00cc"+
		"\u00cd\u0005A\u0000\u0000\u00cd\u00ce\u0005\u001b\u0000\u0000\u00ce\u00cf"+
		"\u0003\u0018\f\u0000\u00cf\u001d\u0001\u0000\u0000\u0000\u00d0\u00d1\u0005"+
		"\u0005\u0000\u0000\u00d1\u00d3\u0005\u001d\u0000\u0000\u00d2\u00d4\u0003"+
		" \u0010\u0000\u00d3\u00d2\u0001\u0000\u0000\u0000\u00d3\u00d4\u0001\u0000"+
		"\u0000\u0000\u00d4\u00d5\u0001\u0000\u0000\u0000\u00d5\u00d6\u0005\u001e"+
		"\u0000\u0000\u00d6\u001f\u0001\u0000\u0000\u0000\u00d7\u00d8\u0003\n\u0005"+
		"\u0000\u00d8\u00de\u0005\u0016\u0000\u0000\u00d9\u00da\u0003\n\u0005\u0000"+
		"\u00da\u00db\u0005\u0016\u0000\u0000\u00db\u00dd\u0001\u0000\u0000\u0000"+
		"\u00dc\u00d9\u0001\u0000\u0000\u0000\u00dd\u00e0\u0001\u0000\u0000\u0000"+
		"\u00de\u00dc\u0001\u0000\u0000\u0000\u00de\u00df\u0001\u0000\u0000\u0000"+
		"\u00df!\u0001\u0000\u0000\u0000\u00e0\u00de\u0001\u0000\u0000\u0000\u00e1"+
		"\u00e6\u0005B\u0000\u0000\u00e2\u00e3\u0005\u001f\u0000\u0000\u00e3\u00e5"+
		"\u0005B\u0000\u0000\u00e4\u00e2\u0001\u0000\u0000\u0000\u00e5\u00e8\u0001"+
		"\u0000\u0000\u0000\u00e6\u00e4\u0001\u0000\u0000\u0000\u00e6\u00e7\u0001"+
		"\u0000\u0000\u0000\u00e7#\u0001\u0000\u0000\u0000\u00e8\u00e6\u0001\u0000"+
		"\u0000\u0000\u00e9\u00ea\u0006\u0012\uffff\uffff\u0000\u00ea\u00ee\u0003"+
		"(\u0014\u0000\u00eb\u00ec\u0007\u0000\u0000\u0000\u00ec\u00ee\u0003$\u0012"+
		"\u0001\u00ed\u00e9\u0001\u0000\u0000\u0000\u00ed\u00eb\u0001\u0000\u0000"+
		"\u0000\u00ee\u00f4\u0001\u0000\u0000\u0000\u00ef\u00f0\n\u0002\u0000\u0000"+
		"\u00f0\u00f1\u0007\u0001\u0000\u0000\u00f1\u00f3\u0003$\u0012\u0003\u00f2"+
		"\u00ef\u0001\u0000\u0000\u0000\u00f3\u00f6\u0001\u0000\u0000\u0000\u00f4"+
		"\u00f2\u0001\u0000\u0000\u0000\u00f4\u00f5\u0001\u0000\u0000\u0000\u00f5"+
		"%\u0001\u0000\u0000\u0000\u00f6\u00f4\u0001\u0000\u0000\u0000\u00f7\u00fc"+
		"\u0003$\u0012\u0000\u00f8\u00f9\u0005\u001f\u0000\u0000\u00f9\u00fb\u0003"+
		"$\u0012\u0000\u00fa\u00f8\u0001\u0000\u0000\u0000\u00fb\u00fe\u0001\u0000"+
		"\u0000\u0000\u00fc\u00fa\u0001\u0000\u0000\u0000\u00fc\u00fd\u0001\u0000"+
		"\u0000\u0000\u00fd\'\u0001\u0000\u0000\u0000\u00fe\u00fc\u0001\u0000\u0000"+
		"\u0000\u00ff\u0100\u0006\u0014\uffff\uffff\u0000\u0100\u0105\u0003*\u0015"+
		"\u0000\u0101\u0105\u00034\u001a\u0000\u0102\u0105\u00036\u001b\u0000\u0103"+
		"\u0105\u00038\u001c\u0000\u0104\u00ff\u0001\u0000\u0000\u0000\u0104\u0101"+
		"\u0001\u0000\u0000\u0000\u0104\u0102\u0001\u0000\u0000\u0000\u0104\u0103"+
		"\u0001\u0000\u0000\u0000\u0105\u010e\u0001\u0000\u0000\u0000\u0106\u0107"+
		"\n\u0006\u0000\u0000\u0107\u010d\u00032\u0019\u0000\u0108\u0109\n\u0005"+
		"\u0000\u0000\u0109\u010d\u0003.\u0017\u0000\u010a\u010b\n\u0004\u0000"+
		"\u0000\u010b\u010d\u00030\u0018\u0000\u010c\u0106\u0001\u0000\u0000\u0000"+
		"\u010c\u0108\u0001\u0000\u0000\u0000\u010c\u010a\u0001\u0000\u0000\u0000"+
		"\u010d\u0110\u0001\u0000\u0000\u0000\u010e\u010c\u0001\u0000\u0000\u0000"+
		"\u010e\u010f\u0001\u0000\u0000\u0000\u010f)\u0001\u0000\u0000\u0000\u0110"+
		"\u010e\u0001\u0000\u0000\u0000\u0111\u0118\u0003,\u0016\u0000\u0112\u0118"+
		"\u0005B\u0000\u0000\u0113\u0114\u0005\u0019\u0000\u0000\u0114\u0115\u0003"+
		"$\u0012\u0000\u0115\u0116\u0005\u001a\u0000\u0000\u0116\u0118\u0001\u0000"+
		"\u0000\u0000\u0117\u0111\u0001\u0000\u0000\u0000\u0117\u0112\u0001\u0000"+
		"\u0000\u0000\u0117\u0113\u0001\u0000\u0000\u0000\u0118+\u0001\u0000\u0000"+
		"\u0000\u0119\u011a\u0007\u0002\u0000\u0000\u011a-\u0001\u0000\u0000\u0000"+
		"\u011b\u011c\u0005\u001c\u0000\u0000\u011c\u011d\u0003$\u0012\u0000\u011d"+
		"\u011e\u0005\u001b\u0000\u0000\u011e/\u0001\u0000\u0000\u0000\u011f\u0121"+
		"\u0005\u0019\u0000\u0000\u0120\u0122\u0003&\u0013\u0000\u0121\u0120\u0001"+
		"\u0000\u0000\u0000\u0121\u0122\u0001\u0000\u0000\u0000\u0122\u0123\u0001"+
		"\u0000\u0000\u0000\u0123\u0124\u0005\u001a\u0000\u0000\u01241\u0001\u0000"+
		"\u0000\u0000\u0125\u0126\u0005\u0018\u0000\u0000\u0126\u0127\u0005B\u0000"+
		"\u0000\u01273\u0001\u0000\u0000\u0000\u0128\u0129\u0005\u0006\u0000\u0000"+
		"\u0129\u012a\u0005\u0019\u0000\u0000\u012a\u012b\u0003$\u0012\u0000\u012b"+
		"\u012c\u0005\u001f\u0000\u0000\u012c\u012d\u0003$\u0012\u0000\u012d\u012e"+
		"\u0005\u001a\u0000\u0000\u012e5\u0001\u0000\u0000\u0000\u012f\u0130\u0005"+
		"\u0007\u0000\u0000\u0130\u0131\u0005\u0019\u0000\u0000\u0131\u0132\u0003"+
		"$\u0012\u0000\u0132\u0133\u0005\u001a\u0000\u0000\u01337\u0001\u0000\u0000"+
		"\u0000\u0134\u0135\u0005\b\u0000\u0000\u0135\u0136\u0005\u0019\u0000\u0000"+
		"\u0136\u0137\u0003$\u0012\u0000\u0137\u0138\u0005\u001a\u0000\u0000\u0138"+
		"9\u0001\u0000\u0000\u0000\u0139\u013b\u0003>\u001f\u0000\u013a\u0139\u0001"+
		"\u0000\u0000\u0000\u013b\u013e\u0001\u0000\u0000\u0000\u013c\u013a\u0001"+
		"\u0000\u0000\u0000\u013c\u013d\u0001\u0000\u0000\u0000\u013d;\u0001\u0000"+
		"\u0000\u0000\u013e\u013c\u0001\u0000\u0000\u0000\u013f\u0140\u0005\u001d"+
		"\u0000\u0000\u0140\u0141\u0003:\u001d\u0000\u0141\u0142\u0005\u001e\u0000"+
		"\u0000\u0142=\u0001\u0000\u0000\u0000\u0143\u0144\u0005\t\u0000\u0000"+
		"\u0144\u0146\u0005\u0019\u0000\u0000\u0145\u0147\u0003&\u0013\u0000\u0146"+
		"\u0145\u0001\u0000\u0000\u0000\u0146\u0147\u0001\u0000\u0000\u0000\u0147"+
		"\u0148\u0001\u0000\u0000\u0000\u0148\u0149\u0005\u001a\u0000\u0000\u0149"+
		"\u016c\u0005\u001f\u0000\u0000\u014a\u014b\u0005\n\u0000\u0000\u014b\u014d"+
		"\u0005\u0019\u0000\u0000\u014c\u014e\u0003&\u0013\u0000\u014d\u014c\u0001"+
		"\u0000\u0000\u0000\u014d\u014e\u0001\u0000\u0000\u0000\u014e\u014f\u0001"+
		"\u0000\u0000\u0000\u014f\u0150\u0005\u001a\u0000\u0000\u0150\u016c\u0005"+
		"\u001f\u0000\u0000\u0151\u0153\u0005\u000b\u0000\u0000\u0152\u0154\u0003"+
		"$\u0012\u0000\u0153\u0152\u0001\u0000\u0000\u0000\u0153\u0154\u0001\u0000"+
		"\u0000\u0000\u0154\u0155\u0001\u0000\u0000\u0000\u0155\u016c\u0005\u001f"+
		"\u0000\u0000\u0156\u0157\u0005\f\u0000\u0000\u0157\u016c\u0005\u001f\u0000"+
		"\u0000\u0158\u0159\u0005\r\u0000\u0000\u0159\u016c\u0005\u001f\u0000\u0000"+
		"\u015a\u015b\u0003@ \u0000\u015b\u015c\u0005\u001f\u0000\u0000\u015c\u016c"+
		"\u0001\u0000\u0000\u0000\u015d\u015e\u0003<\u001e\u0000\u015e\u015f\u0005"+
		"\u001f\u0000\u0000\u015f\u016c\u0001\u0000\u0000\u0000\u0160\u0161\u0003"+
		"H$\u0000\u0161\u0162\u0005\u001f\u0000\u0000\u0162\u016c\u0001\u0000\u0000"+
		"\u0000\u0163\u0164\u0003D\"\u0000\u0164\u0165\u0005\u001f\u0000\u0000"+
		"\u0165\u016c\u0001\u0000\u0000\u0000\u0166\u0167\u0003F#\u0000\u0167\u0168"+
		"\u0005\u001f\u0000\u0000\u0168\u016c\u0001\u0000\u0000\u0000\u0169\u016c"+
		"\u0003\f\u0006\u0000\u016a\u016c\u0003\u0004\u0002\u0000\u016b\u0143\u0001"+
		"\u0000\u0000\u0000\u016b\u014a\u0001\u0000\u0000\u0000\u016b\u0151\u0001"+
		"\u0000\u0000\u0000\u016b\u0156\u0001\u0000\u0000\u0000\u016b\u0158\u0001"+
		"\u0000\u0000\u0000\u016b\u015a\u0001\u0000\u0000\u0000\u016b\u015d\u0001"+
		"\u0000\u0000\u0000\u016b\u0160\u0001\u0000\u0000\u0000\u016b\u0163\u0001"+
		"\u0000\u0000\u0000\u016b\u0166\u0001\u0000\u0000\u0000\u016b\u0169\u0001"+
		"\u0000\u0000\u0000\u016b\u016a\u0001\u0000\u0000\u0000\u016c?\u0001\u0000"+
		"\u0000\u0000\u016d\u016f\u0003$\u0012\u0000\u016e\u0170\u0007\u0003\u0000"+
		"\u0000\u016f\u016e\u0001\u0000\u0000\u0000\u016f\u0170\u0001\u0000\u0000"+
		"\u0000\u0170\u0177\u0001\u0000\u0000\u0000\u0171\u0177\u0003B!\u0000\u0172"+
		"\u0173\u0003&\u0013\u0000\u0173\u0174\u0005\u0015\u0000\u0000\u0174\u0175"+
		"\u0003&\u0013\u0000\u0175\u0177\u0001\u0000\u0000\u0000\u0176\u016d\u0001"+
		"\u0000\u0000\u0000\u0176\u0171\u0001\u0000\u0000\u0000\u0176\u0172\u0001"+
		"\u0000\u0000\u0000\u0177A\u0001\u0000\u0000\u0000\u0178\u0179\u0003&\u0013"+
		"\u0000\u0179\u017a\u0005\u0014\u0000\u0000\u017a\u017b\u0003&\u0013\u0000"+
		"\u017b\u0181\u0001\u0000\u0000\u0000\u017c\u017d\u0003$\u0012\u0000\u017d"+
		"\u017e\u0007\u0004\u0000\u0000\u017e\u017f\u0003$\u0012\u0000\u017f\u0181"+
		"\u0001\u0000\u0000\u0000\u0180\u0178\u0001\u0000\u0000\u0000\u0180\u017c"+
		"\u0001\u0000\u0000\u0000\u0181C\u0001\u0000\u0000\u0000\u0182\u0183\u0005"+
		"\u000e\u0000\u0000\u0183\u0184\u0003$\u0012\u0000\u0184\u0185\u0003<\u001e"+
		"\u0000\u0185\u01a9\u0001\u0000\u0000\u0000\u0186\u0187\u0005\u000e\u0000"+
		"\u0000\u0187\u0188\u0003$\u0012\u0000\u0188\u0189\u0003<\u001e\u0000\u0189"+
		"\u018a\u0005\u000f\u0000\u0000\u018a\u018b\u0003D\"\u0000\u018b\u01a9"+
		"\u0001\u0000\u0000\u0000\u018c\u018d\u0005\u000e\u0000\u0000\u018d\u018e"+
		"\u0003$\u0012\u0000\u018e\u018f\u0003<\u001e\u0000\u018f\u0190\u0005\u000f"+
		"\u0000\u0000\u0190\u0191\u0003<\u001e\u0000\u0191\u01a9\u0001\u0000\u0000"+
		"\u0000\u0192\u0193\u0005\u000e\u0000\u0000\u0193\u0194\u0003@ \u0000\u0194"+
		"\u0195\u0005\u0016\u0000\u0000\u0195\u0196\u0003$\u0012\u0000\u0196\u0197"+
		"\u0003<\u001e\u0000\u0197\u01a9\u0001\u0000\u0000\u0000\u0198\u0199\u0005"+
		"\u000e\u0000\u0000\u0199\u019a\u0003@ \u0000\u019a\u019b\u0005\u0016\u0000"+
		"\u0000\u019b\u019c\u0003$\u0012\u0000\u019c\u019d\u0003<\u001e\u0000\u019d"+
		"\u019e\u0005\u000f\u0000\u0000\u019e\u019f\u0003D\"\u0000\u019f\u01a9"+
		"\u0001\u0000\u0000\u0000\u01a0\u01a1\u0005\u000e\u0000\u0000\u01a1\u01a2"+
		"\u0003@ \u0000\u01a2\u01a3\u0005\u0016\u0000\u0000\u01a3\u01a4\u0003$"+
		"\u0012\u0000\u01a4\u01a5\u0003<\u001e\u0000\u01a5\u01a6\u0005\u000f\u0000"+
		"\u0000\u01a6\u01a7\u0003<\u001e\u0000\u01a7\u01a9\u0001\u0000\u0000\u0000"+
		"\u01a8\u0182\u0001\u0000\u0000\u0000\u01a8\u0186\u0001\u0000\u0000\u0000"+
		"\u01a8\u018c\u0001\u0000\u0000\u0000\u01a8\u0192\u0001\u0000\u0000\u0000"+
		"\u01a8\u0198\u0001\u0000\u0000\u0000\u01a8\u01a0\u0001\u0000\u0000\u0000"+
		"\u01a9E\u0001\u0000\u0000\u0000\u01aa\u01ab\u0005\u0010\u0000\u0000\u01ab"+
		"\u01c0\u0003<\u001e\u0000\u01ac\u01ad\u0005\u0010\u0000\u0000\u01ad\u01ae"+
		"\u0003$\u0012\u0000\u01ae\u01af\u0003<\u001e\u0000\u01af\u01c0\u0001\u0000"+
		"\u0000\u0000\u01b0\u01b1\u0005\u0010\u0000\u0000\u01b1\u01b2\u0003@ \u0000"+
		"\u01b2\u01b3\u0005\u0016\u0000\u0000\u01b3\u01b4\u0003$\u0012\u0000\u01b4"+
		"\u01b5\u0005\u0016\u0000\u0000\u01b5\u01b6\u0003@ \u0000\u01b6\u01b7\u0003"+
		"<\u001e\u0000\u01b7\u01c0\u0001\u0000\u0000\u0000\u01b8\u01b9\u0005\u0010"+
		"\u0000\u0000\u01b9\u01ba\u0003@ \u0000\u01ba\u01bb\u0005\u0016\u0000\u0000"+
		"\u01bb\u01bc\u0005\u0016\u0000\u0000\u01bc\u01bd\u0003@ \u0000\u01bd\u01be"+
		"\u0003<\u001e\u0000\u01be\u01c0\u0001\u0000\u0000\u0000\u01bf\u01aa\u0001"+
		"\u0000\u0000\u0000\u01bf\u01ac\u0001\u0000\u0000\u0000\u01bf\u01b0\u0001"+
		"\u0000\u0000\u0000\u01bf\u01b8\u0001\u0000\u0000\u0000\u01c0G\u0001\u0000"+
		"\u0000\u0000\u01c1\u01c3\u0005\u0011\u0000\u0000\u01c2\u01c4\u0003@ \u0000"+
		"\u01c3\u01c2\u0001\u0000\u0000\u0000\u01c3\u01c4\u0001\u0000\u0000\u0000"+
		"\u01c4\u01c6\u0001\u0000\u0000\u0000\u01c5\u01c7\u0005\u0016\u0000\u0000"+
		"\u01c6\u01c5\u0001\u0000\u0000\u0000\u01c6\u01c7\u0001\u0000\u0000\u0000"+
		"\u01c7\u01c9\u0001\u0000\u0000\u0000\u01c8\u01ca\u0003$\u0012\u0000\u01c9"+
		"\u01c8\u0001\u0000\u0000\u0000\u01c9\u01ca\u0001\u0000\u0000\u0000\u01ca"+
		"\u01cb\u0001\u0000\u0000\u0000\u01cb\u01cc\u0005\u001d\u0000\u0000\u01cc"+
		"\u01cd\u0003J%\u0000\u01cd\u01ce\u0005\u001e\u0000\u0000\u01ceI\u0001"+
		"\u0000\u0000\u0000\u01cf\u01d1\u0003L&\u0000\u01d0\u01cf\u0001\u0000\u0000"+
		"\u0000\u01d1\u01d4\u0001\u0000\u0000\u0000\u01d2\u01d0\u0001\u0000\u0000"+
		"\u0000\u01d2\u01d3\u0001\u0000\u0000\u0000\u01d3K\u0001\u0000\u0000\u0000"+
		"\u01d4\u01d2\u0001\u0000\u0000\u0000\u01d5\u01d6\u0003N\'\u0000\u01d6"+
		"\u01d7\u0005\u0017\u0000\u0000\u01d7\u01d8\u0003:\u001d\u0000\u01d8M\u0001"+
		"\u0000\u0000\u0000\u01d9\u01da\u0005\u0012\u0000\u0000\u01da\u01dd\u0003"+
		"&\u0013\u0000\u01db\u01dd\u0005\u0013\u0000\u0000\u01dc\u01d9\u0001\u0000"+
		"\u0000\u0000\u01dc\u01db\u0001\u0000\u0000\u0000\u01ddO\u0001\u0000\u0000"+
		"\u0000&XZadmy\u0082\u0085\u0090\u00b0\u00b7\u00c5\u00d3\u00de\u00e6\u00ed"+
		"\u00f4\u00fc\u0104\u010c\u010e\u0117\u0121\u013c\u0146\u014d\u0153\u016b"+
		"\u016f\u0176\u0180\u01a8\u01bf\u01c3\u01c6\u01c9\u01d2\u01dc";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}