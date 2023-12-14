// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package syslog

import (
	"bytes"
	"context"
	"fmt"
	"strconv"
<<<<<<< HEAD
	"strings"
=======
>>>>>>> 4cb759cfc9 (fixed log)
	"sync"

	"github.com/hashicorp/eventlogger"
	gsyslog "github.com/hashicorp/go-syslog"
	"github.com/hashicorp/vault/audit"
	"github.com/hashicorp/vault/internal/observability/event"
	"github.com/hashicorp/vault/sdk/helper/salt"
	"github.com/hashicorp/vault/sdk/logical"
)

<<<<<<< HEAD
var _ audit.Backend = (*Backend)(nil)

// Backend is the audit backend for the syslog-based audit store.
type Backend struct {
	formatter    *audit.EntryFormatterWriter
	formatConfig audit.FormatterConfig
	logger       gsyslog.Syslogger
	name         string
	nodeIDList   []eventlogger.NodeID
	nodeMap      map[eventlogger.NodeID]eventlogger.Node
	salt         *salt.Salt
	saltConfig   *salt.Config
	saltMutex    sync.RWMutex
	saltView     logical.Storage
}

func Factory(_ context.Context, conf *audit.BackendConfig, useEventLogger bool, headersConfig audit.HeaderFormatter) (audit.Backend, error) {
	const op = "syslog.Factory"

	if conf.SaltConfig == nil {
		return nil, fmt.Errorf("%s: nil salt config", op)
	}

	if conf.SaltView == nil {
		return nil, fmt.Errorf("%s: nil salt view", op)
=======
func Factory(ctx context.Context, conf *audit.BackendConfig, useEventLogger bool, headersConfig audit.HeaderFormatter) (audit.Backend, error) {
	if conf.SaltConfig == nil {
		return nil, fmt.Errorf("nil salt config")
	}

	if conf.SaltView == nil {
		return nil, fmt.Errorf("nil salt view")
>>>>>>> 4cb759cfc9 (fixed log)
	}

	// Get facility or default to AUTH
	facility, ok := conf.Config["facility"]
	if !ok {
		facility = "AUTH"
	}

	// Get tag or default to 'vault'
	tag, ok := conf.Config["tag"]
	if !ok {
		tag = "vault"
	}

<<<<<<< HEAD
	cfg, err := formatterConfig(conf.Config)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to create formatter config: %w", op, err)
=======
	var cfgOpts []audit.Option

	if format, ok := conf.Config["format"]; ok {
		cfgOpts = append(cfgOpts, audit.WithFormat(format))
	}

	// Check if hashing of accessor is disabled
	if hmacAccessorRaw, ok := conf.Config["hmac_accessor"]; ok {
		v, err := strconv.ParseBool(hmacAccessorRaw)
		if err != nil {
			return nil, err
		}
		cfgOpts = append(cfgOpts, audit.WithHMACAccessor(v))
	}

	// Check if raw logging is enabled
	if raw, ok := conf.Config["log_raw"]; ok {
		v, err := strconv.ParseBool(raw)
		if err != nil {
			return nil, err
		}
		cfgOpts = append(cfgOpts, audit.WithRaw(v))
	}

	if elideListResponsesRaw, ok := conf.Config["elide_list_responses"]; ok {
		v, err := strconv.ParseBool(elideListResponsesRaw)
		if err != nil {
			return nil, err
		}
		cfgOpts = append(cfgOpts, audit.WithElision(v))
	}

	cfg, err := audit.NewFormatterConfig(cfgOpts...)
	if err != nil {
		return nil, err
>>>>>>> 4cb759cfc9 (fixed log)
	}

	// Get the logger
	logger, err := gsyslog.NewLogger(gsyslog.LOG_INFO, facility, tag)
	if err != nil {
<<<<<<< HEAD
		return nil, fmt.Errorf("%s: cannot create logger: %w", op, err)
	}

	b := &Backend{
		formatConfig: cfg,
		logger:       logger,
		name:         conf.MountPath,
		saltConfig:   conf.SaltConfig,
		saltView:     conf.SaltView,
=======
		return nil, err
	}

	b := &Backend{
		logger:       logger,
		saltConfig:   conf.SaltConfig,
		saltView:     conf.SaltView,
		formatConfig: cfg,
>>>>>>> 4cb759cfc9 (fixed log)
	}

	// Configure the formatter for either case.
	f, err := audit.NewEntryFormatter(b.formatConfig, b, audit.WithHeaderFormatter(headersConfig), audit.WithPrefix(conf.Config["prefix"]))
	if err != nil {
<<<<<<< HEAD
		return nil, fmt.Errorf("%s: error creating formatter: %w", op, err)
=======
		return nil, fmt.Errorf("error creating formatter: %w", err)
>>>>>>> 4cb759cfc9 (fixed log)
	}

	var w audit.Writer
	switch b.formatConfig.RequiredFormat {
	case audit.JSONFormat:
		w = &audit.JSONWriter{Prefix: conf.Config["prefix"]}
	case audit.JSONxFormat:
		w = &audit.JSONxWriter{Prefix: conf.Config["prefix"]}
	}

	fw, err := audit.NewEntryFormatterWriter(b.formatConfig, f, w)
	if err != nil {
<<<<<<< HEAD
		return nil, fmt.Errorf("%s: error creating formatter writer: %w", op, err)
=======
		return nil, fmt.Errorf("error creating formatter writer: %w", err)
>>>>>>> 4cb759cfc9 (fixed log)
	}

	b.formatter = fw

	if useEventLogger {
<<<<<<< HEAD
		b.nodeIDList = []eventlogger.NodeID{}
		b.nodeMap = make(map[eventlogger.NodeID]eventlogger.Node)

		err := b.configureFilterNode(conf.Config["filter"])
		if err != nil {
			return nil, fmt.Errorf("%s: error configuring filter node: %w", op, err)
		}

		formatterOpts := []audit.Option{
			audit.WithHeaderFormatter(headersConfig),
			audit.WithPrefix(conf.Config["prefix"]),
		}

		err = b.configureFormatterNode(cfg, formatterOpts...)
		if err != nil {
			return nil, fmt.Errorf("%s: error configuring formatter node: %w", op, err)
		}

		sinkOpts := []event.Option{
			event.WithFacility(facility),
			event.WithTag(tag),
		}

		err = b.configureSinkNode(conf.MountPath, cfg.RequiredFormat.String(), sinkOpts...)
		if err != nil {
			return nil, fmt.Errorf("%s: error configuring sink node: %w", op, err)
		}
	}

	return b, nil
}

// Deprecated: Use eventlogger.
=======
		var opts []event.Option

		// Get facility or default to AUTH
		if facility, ok := conf.Config["facility"]; ok {
			opts = append(opts, event.WithFacility(facility))
		}

		if tag, ok := conf.Config["tag"]; ok {
			opts = append(opts, event.WithTag(tag))
		}

		b.nodeIDList = make([]eventlogger.NodeID, 2)
		b.nodeMap = make(map[eventlogger.NodeID]eventlogger.Node)

		formatterNodeID, err := event.GenerateNodeID()
		if err != nil {
			return nil, fmt.Errorf("error generating random NodeID for formatter node: %w", err)
		}
		b.nodeIDList[0] = formatterNodeID
		b.nodeMap[formatterNodeID] = f

		n, err := event.NewSyslogSink(b.formatConfig.RequiredFormat.String(), opts...)
		if err != nil {
			return nil, fmt.Errorf("error creating syslog sink node: %w", err)
		}
		sinkNode := &audit.SinkWrapper{Name: conf.MountPath, Sink: n}

		sinkNodeID, err := event.GenerateNodeID()
		if err != nil {
			return nil, fmt.Errorf("error generating random NodeID for sink node: %w", err)
		}
		b.nodeIDList[1] = sinkNodeID
		b.nodeMap[sinkNodeID] = sinkNode
	}
	return b, nil
}

// Backend is the audit backend for the syslog-based audit store.
type Backend struct {
	logger gsyslog.Syslogger

	formatter    *audit.EntryFormatterWriter
	formatConfig audit.FormatterConfig

	saltMutex  sync.RWMutex
	salt       *salt.Salt
	saltConfig *salt.Config
	saltView   logical.Storage

	nodeIDList []eventlogger.NodeID
	nodeMap    map[eventlogger.NodeID]eventlogger.Node
}

var _ audit.Backend = (*Backend)(nil)

>>>>>>> 4cb759cfc9 (fixed log)
func (b *Backend) LogRequest(ctx context.Context, in *logical.LogInput) error {
	var buf bytes.Buffer
	if err := b.formatter.FormatAndWriteRequest(ctx, &buf, in); err != nil {
		return err
	}

	// Write out to syslog
	_, err := b.logger.Write(buf.Bytes())
	return err
}

<<<<<<< HEAD
// Deprecated: Use eventlogger.
=======
>>>>>>> 4cb759cfc9 (fixed log)
func (b *Backend) LogResponse(ctx context.Context, in *logical.LogInput) error {
	var buf bytes.Buffer
	if err := b.formatter.FormatAndWriteResponse(ctx, &buf, in); err != nil {
		return err
	}

	// Write out to syslog
	_, err := b.logger.Write(buf.Bytes())
	return err
}

func (b *Backend) LogTestMessage(ctx context.Context, in *logical.LogInput, config map[string]string) error {
	// Event logger behavior - manually Process each node
	if len(b.nodeIDList) > 0 {
		return audit.ProcessManual(ctx, in, b.nodeIDList, b.nodeMap)
	}

	// Old behavior
	var buf bytes.Buffer

	temporaryFormatter, err := audit.NewTemporaryFormatter(config["format"], config["prefix"])
	if err != nil {
		return err
	}

	if err = temporaryFormatter.FormatAndWriteRequest(ctx, &buf, in); err != nil {
		return err
	}

	// Send to syslog
	_, err = b.logger.Write(buf.Bytes())
	return err
}

func (b *Backend) Reload(_ context.Context) error {
	return nil
}

func (b *Backend) Salt(ctx context.Context) (*salt.Salt, error) {
	b.saltMutex.RLock()
	if b.salt != nil {
		defer b.saltMutex.RUnlock()
		return b.salt, nil
	}
	b.saltMutex.RUnlock()
	b.saltMutex.Lock()
	defer b.saltMutex.Unlock()
	if b.salt != nil {
		return b.salt, nil
	}
<<<<<<< HEAD
	s, err := salt.NewSalt(ctx, b.saltView, b.saltConfig)
	if err != nil {
		return nil, err
	}
	b.salt = s
	return s, nil
=======
	salt, err := salt.NewSalt(ctx, b.saltView, b.saltConfig)
	if err != nil {
		return nil, err
	}
	b.salt = salt
	return salt, nil
>>>>>>> 4cb759cfc9 (fixed log)
}

func (b *Backend) Invalidate(_ context.Context) {
	b.saltMutex.Lock()
	defer b.saltMutex.Unlock()
	b.salt = nil
}

<<<<<<< HEAD
// formatterConfig creates the configuration required by a formatter node using
// the config map supplied to the factory.
func formatterConfig(config map[string]string) (audit.FormatterConfig, error) {
	const op = "syslog.formatterConfig"

	var opts []audit.Option

	if format, ok := config["format"]; ok {
		opts = append(opts, audit.WithFormat(format))
	}

	// Check if hashing of accessor is disabled
	if hmacAccessorRaw, ok := config["hmac_accessor"]; ok {
		v, err := strconv.ParseBool(hmacAccessorRaw)
		if err != nil {
			return audit.FormatterConfig{}, fmt.Errorf("%s: unable to parse 'hmac_accessor': %w", op, err)
		}
		opts = append(opts, audit.WithHMACAccessor(v))
	}

	// Check if raw logging is enabled
	if raw, ok := config["log_raw"]; ok {
		v, err := strconv.ParseBool(raw)
		if err != nil {
			return audit.FormatterConfig{}, fmt.Errorf("%s: unable to parse 'log_raw': %w", op, err)
		}
		opts = append(opts, audit.WithRaw(v))
	}

	if elideListResponsesRaw, ok := config["elide_list_responses"]; ok {
		v, err := strconv.ParseBool(elideListResponsesRaw)
		if err != nil {
			return audit.FormatterConfig{}, fmt.Errorf("%s: unable to parse 'elide_list_responses': %w", op, err)
		}
		opts = append(opts, audit.WithElision(v))
	}

	return audit.NewFormatterConfig(opts...)
}

// configureFilterNode is used to configure a filter node and associated ID on the Backend.
func (b *Backend) configureFilterNode(filter string) error {
	const op = "syslog.(Backend).configureFilterNode"

	filter = strings.TrimSpace(filter)
	if filter == "" {
		return nil
	}

	filterNodeID, err := event.GenerateNodeID()
	if err != nil {
		return fmt.Errorf("%s: error generating random NodeID for filter node: %w", op, err)
	}

	filterNode, err := audit.NewEntryFilter(filter)
	if err != nil {
		return fmt.Errorf("%s: error creating filter node: %w", op, err)
	}

	b.nodeIDList = append(b.nodeIDList, filterNodeID)
	b.nodeMap[filterNodeID] = filterNode
	return nil
}

// configureFormatterNode is used to configure a formatter node and associated ID on the Backend.
func (b *Backend) configureFormatterNode(formatConfig audit.FormatterConfig, opts ...audit.Option) error {
	const op = "syslog.(Backend).configureFormatterNode"

	formatterNodeID, err := event.GenerateNodeID()
	if err != nil {
		return fmt.Errorf("%s: error generating random NodeID for formatter node: %w", op, err)
	}

	formatterNode, err := audit.NewEntryFormatter(formatConfig, b, opts...)
	if err != nil {
		return fmt.Errorf("%s: error creating formatter: %w", op, err)
	}

	b.nodeIDList = append(b.nodeIDList, formatterNodeID)
	b.nodeMap[formatterNodeID] = formatterNode
	return nil
}

// configureSinkNode is used to configure a sink node and associated ID on the Backend.
func (b *Backend) configureSinkNode(name string, format string, opts ...event.Option) error {
	const op = "syslog.(Backend).configureSinkNode"

	name = strings.TrimSpace(name)
	if name == "" {
		return fmt.Errorf("%s: name is required: %w", op, event.ErrInvalidParameter)
	}

	format = strings.TrimSpace(format)
	if format == "" {
		return fmt.Errorf("%s: format is required: %w", op, event.ErrInvalidParameter)
	}

	sinkNodeID, err := event.GenerateNodeID()
	if err != nil {
		return fmt.Errorf("%s: error generating random NodeID for sink node: %w", op, err)
	}

	n, err := event.NewSyslogSink(format, opts...)
	if err != nil {
		return fmt.Errorf("%s: error creating syslog sink node: %w", op, err)
	}

	// wrap the sink node with metrics middleware
	sinkNode := &audit.SinkWrapper{Name: name, Sink: n}

	b.nodeIDList = append(b.nodeIDList, sinkNodeID)
	b.nodeMap[sinkNodeID] = sinkNode
	return nil
}

// Name for this backend, this would ideally correspond to the mount path for the audit device.
func (b *Backend) Name() string {
	return b.name
}

// Nodes returns the nodes which should be used by the event framework to process audit entries.
func (b *Backend) Nodes() map[eventlogger.NodeID]eventlogger.Node {
	return b.nodeMap
}

// NodeIDs returns the IDs of the nodes, in the order they are required.
func (b *Backend) NodeIDs() []eventlogger.NodeID {
	return b.nodeIDList
}

// EventType returns the event type for the backend.
func (b *Backend) EventType() eventlogger.EventType {
	return eventlogger.EventType(event.AuditType.String())
}

// HasFiltering determines if the first node for the pipeline is an eventlogger.NodeTypeFilter.
func (b *Backend) HasFiltering() bool {
	return len(b.nodeIDList) > 0 && b.nodeMap[b.nodeIDList[0]].Type() == eventlogger.NodeTypeFilter
=======
// RegisterNodesAndPipeline registers the nodes and a pipeline as required by
// the audit.Backend interface.
func (b *Backend) RegisterNodesAndPipeline(broker *eventlogger.Broker, name string) error {
	for id, node := range b.nodeMap {
		if err := broker.RegisterNode(id, node); err != nil {
			return err
		}
	}

	pipeline := eventlogger.Pipeline{
		PipelineID: eventlogger.PipelineID(name),
		EventType:  eventlogger.EventType("audit"),
		NodeIDs:    b.nodeIDList,
	}

	return broker.RegisterPipeline(pipeline)
>>>>>>> 4cb759cfc9 (fixed log)
}
