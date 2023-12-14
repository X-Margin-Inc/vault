// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package file

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/hashicorp/eventlogger"
	"github.com/hashicorp/vault/audit"
	"github.com/hashicorp/vault/internal/observability/event"
	"github.com/hashicorp/vault/sdk/helper/salt"
	"github.com/hashicorp/vault/sdk/logical"
)

const (
	stdout  = "stdout"
	discard = "discard"
)

<<<<<<< HEAD
var _ audit.Backend = (*Backend)(nil)

// Backend is the audit backend for the file-based audit store.
//
// NOTE: This audit backend is currently very simple: it appends to a file.
// It doesn't do anything more at the moment to assist with rotation
// or reset the write cursor, this should be done in the future.
type Backend struct {
	f            *os.File
	fileLock     sync.RWMutex
	formatter    *audit.EntryFormatterWriter
	formatConfig audit.FormatterConfig
	mode         os.FileMode
	name         string
	nodeIDList   []eventlogger.NodeID
	nodeMap      map[eventlogger.NodeID]eventlogger.Node
	filePath     string
	salt         *atomic.Value
	saltConfig   *salt.Config
	saltMutex    sync.RWMutex
	saltView     logical.Storage
}

func Factory(_ context.Context, conf *audit.BackendConfig, useEventLogger bool, headersConfig audit.HeaderFormatter) (audit.Backend, error) {
	const op = "file.Factory"

	if conf.SaltConfig == nil {
		return nil, fmt.Errorf("%s: nil salt config", op)
	}
	if conf.SaltView == nil {
		return nil, fmt.Errorf("%s: nil salt view", op)
	}

	// Get file path from config or fall back to the old option name ('path') for compatibility
	// (see commit bac4fe0799a372ba1245db642f3f6cd1f1d02669).
	var filePath string
	if p, ok := conf.Config["file_path"]; ok {
		filePath = p
	} else if p, ok = conf.Config["path"]; ok {
		filePath = p
	} else {
		return nil, fmt.Errorf("%s: file_path is required", op)
	}

	// normalize file path if configured for stdout
	if strings.EqualFold(filePath, stdout) {
		filePath = stdout
	}
	if strings.EqualFold(filePath, discard) {
		filePath = discard
=======
func Factory(ctx context.Context, conf *audit.BackendConfig, useEventLogger bool, headersConfig audit.HeaderFormatter) (audit.Backend, error) {
	if conf.SaltConfig == nil {
		return nil, fmt.Errorf("nil salt config")
	}
	if conf.SaltView == nil {
		return nil, fmt.Errorf("nil salt view")
	}

	path, ok := conf.Config["file_path"]
	if !ok {
		path, ok = conf.Config["path"]
		if !ok {
			return nil, fmt.Errorf("file_path is required")
		}
	}

	// normalize path if configured for stdout
	if strings.EqualFold(path, stdout) {
		path = stdout
	}
	if strings.EqualFold(path, discard) {
		path = discard
	}

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
>>>>>>> 4cb759cfc9 (fixed log)
	}

	mode := os.FileMode(0o600)
	if modeRaw, ok := conf.Config["mode"]; ok {
		m, err := strconv.ParseUint(modeRaw, 8, 32)
		if err != nil {
<<<<<<< HEAD
			return nil, fmt.Errorf("%s: unable to parse 'mode': %w", op, err)
=======
			return nil, err
>>>>>>> 4cb759cfc9 (fixed log)
		}
		switch m {
		case 0:
			// if mode is 0000, then do not modify file mode
<<<<<<< HEAD
			if filePath != stdout && filePath != discard {
				fileInfo, err := os.Stat(filePath)
				if err != nil {
					return nil, fmt.Errorf("%s: unable to stat %q: %w", op, filePath, err)
=======
			if path != stdout && path != discard {
				fileInfo, err := os.Stat(path)
				if err != nil {
					return nil, err
>>>>>>> 4cb759cfc9 (fixed log)
				}
				mode = fileInfo.Mode()
			}
		default:
			mode = os.FileMode(m)
		}
	}

<<<<<<< HEAD
	cfg, err := formatterConfig(conf.Config)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to create formatter config: %w", op, err)
	}

	b := &Backend{
		filePath:     filePath,
		formatConfig: cfg,
		mode:         mode,
		name:         conf.MountPath,
		saltConfig:   conf.SaltConfig,
		saltView:     conf.SaltView,
		salt:         new(atomic.Value),
=======
	cfg, err := audit.NewFormatterConfig(cfgOpts...)
	if err != nil {
		return nil, err
	}

	b := &Backend{
		path:         path,
		mode:         mode,
		saltConfig:   conf.SaltConfig,
		saltView:     conf.SaltView,
		salt:         new(atomic.Value),
		formatConfig: cfg,
>>>>>>> 4cb759cfc9 (fixed log)
	}

	// Ensure we are working with the right type by explicitly storing a nil of
	// the right type
	b.salt.Store((*salt.Salt)(nil))

	// Configure the formatter for either case.
	f, err := audit.NewEntryFormatter(b.formatConfig, b, audit.WithHeaderFormatter(headersConfig), audit.WithPrefix(conf.Config["prefix"]))
	if err != nil {
<<<<<<< HEAD
		return nil, fmt.Errorf("%s: error creating formatter: %w", op, err)
	}

=======
		return nil, fmt.Errorf("error creating formatter: %w", err)
	}
>>>>>>> 4cb759cfc9 (fixed log)
	var w audit.Writer
	switch b.formatConfig.RequiredFormat {
	case audit.JSONFormat:
		w = &audit.JSONWriter{Prefix: conf.Config["prefix"]}
	case audit.JSONxFormat:
		w = &audit.JSONxWriter{Prefix: conf.Config["prefix"]}
	default:
<<<<<<< HEAD
		return nil, fmt.Errorf("%s: unknown format type %q", op, b.formatConfig.RequiredFormat)
=======
		return nil, fmt.Errorf("unknown format type %q", b.formatConfig.RequiredFormat)
>>>>>>> 4cb759cfc9 (fixed log)
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

		err = b.configureSinkNode(conf.MountPath, filePath, conf.Config["mode"], cfg.RequiredFormat.String())
		if err != nil {
			return nil, fmt.Errorf("%s: error configuring sink node: %w", op, err)
		}
	} else {
		switch filePath {
=======
		b.nodeIDList = make([]eventlogger.NodeID, 2)
		b.nodeMap = make(map[eventlogger.NodeID]eventlogger.Node)

		formatterNodeID, err := event.GenerateNodeID()
		if err != nil {
			return nil, fmt.Errorf("error generating random NodeID for formatter node: %w", err)
		}

		b.nodeIDList[0] = formatterNodeID
		b.nodeMap[formatterNodeID] = f

		var sinkNode eventlogger.Node

		switch path {
		case stdout:
			sinkNode = &audit.SinkWrapper{Name: path, Sink: event.NewStdoutSinkNode(b.formatConfig.RequiredFormat.String())}
		case discard:
			sinkNode = &audit.SinkWrapper{Name: path, Sink: event.NewNoopSink()}
		default:
			var err error

			var opts []event.Option
			// Check if mode is provided
			if modeRaw, ok := conf.Config["mode"]; ok {
				opts = append(opts, event.WithFileMode(modeRaw))
			}

			// The NewFileSink function attempts to open the file and will
			// return an error if it can't.
			n, err := event.NewFileSink(
				b.path,
				b.formatConfig.RequiredFormat.String(), opts...)
			if err != nil {
				return nil, fmt.Errorf("file sink creation failed for path %q: %w", path, err)
			}
			sinkNode = &audit.SinkWrapper{Name: conf.MountPath, Sink: n}
		}

		sinkNodeID, err := event.GenerateNodeID()
		if err != nil {
			return nil, fmt.Errorf("error generating random NodeID for sink node: %w", err)
		}

		b.nodeIDList[1] = sinkNodeID
		b.nodeMap[sinkNodeID] = sinkNode
	} else {
		switch path {
>>>>>>> 4cb759cfc9 (fixed log)
		case stdout:
		case discard:
		default:
			// Ensure that the file can be successfully opened for writing;
			// otherwise it will be too late to catch later without problems
			// (ref: https://github.com/hashicorp/vault/issues/550)
			if err := b.open(); err != nil {
<<<<<<< HEAD
				return nil, fmt.Errorf("%s: sanity check failed; unable to open %q for writing: %w", op, filePath, err)
=======
				return nil, fmt.Errorf("sanity check failed; unable to open %q for writing: %w", path, err)
>>>>>>> 4cb759cfc9 (fixed log)
			}
		}
	}

	return b, nil
}

<<<<<<< HEAD
=======
// Backend is the audit backend for the file-based audit store.
//
// NOTE: This audit backend is currently very simple: it appends to a file.
// It doesn't do anything more at the moment to assist with rotation
// or reset the write cursor, this should be done in the future.
type Backend struct {
	path string

	formatter    *audit.EntryFormatterWriter
	formatConfig audit.FormatterConfig

	fileLock sync.RWMutex
	f        *os.File
	mode     os.FileMode

	saltMutex  sync.RWMutex
	salt       *atomic.Value
	saltConfig *salt.Config
	saltView   logical.Storage

	nodeIDList []eventlogger.NodeID
	nodeMap    map[eventlogger.NodeID]eventlogger.Node
}

var _ audit.Backend = (*Backend)(nil)

>>>>>>> 4cb759cfc9 (fixed log)
func (b *Backend) Salt(ctx context.Context) (*salt.Salt, error) {
	s := b.salt.Load().(*salt.Salt)
	if s != nil {
		return s, nil
	}

	b.saltMutex.Lock()
	defer b.saltMutex.Unlock()

	s = b.salt.Load().(*salt.Salt)
	if s != nil {
		return s, nil
	}

	newSalt, err := salt.NewSalt(ctx, b.saltView, b.saltConfig)
	if err != nil {
		b.salt.Store((*salt.Salt)(nil))
		return nil, err
	}

	b.salt.Store(newSalt)
	return newSalt, nil
}

<<<<<<< HEAD
// Deprecated: Use eventlogger.
func (b *Backend) LogRequest(ctx context.Context, in *logical.LogInput) error {
	var writer io.Writer
	switch b.filePath {
=======
func (b *Backend) LogRequest(ctx context.Context, in *logical.LogInput) error {
	var writer io.Writer
	switch b.path {
>>>>>>> 4cb759cfc9 (fixed log)
	case stdout:
		writer = os.Stdout
	case discard:
		return nil
	}

	buf := bytes.NewBuffer(make([]byte, 0, 2000))
	err := b.formatter.FormatAndWriteRequest(ctx, buf, in)
	if err != nil {
		return err
	}

	return b.log(ctx, buf, writer)
}

<<<<<<< HEAD
// Deprecated: Use eventlogger.
=======
>>>>>>> 4cb759cfc9 (fixed log)
func (b *Backend) log(_ context.Context, buf *bytes.Buffer, writer io.Writer) error {
	reader := bytes.NewReader(buf.Bytes())

	b.fileLock.Lock()

	if writer == nil {
		if err := b.open(); err != nil {
			b.fileLock.Unlock()
			return err
		}
		writer = b.f
	}

	if _, err := reader.WriteTo(writer); err == nil {
		b.fileLock.Unlock()
		return nil
<<<<<<< HEAD
	} else if b.filePath == stdout {
=======
	} else if b.path == stdout {
>>>>>>> 4cb759cfc9 (fixed log)
		b.fileLock.Unlock()
		return err
	}

	// If writing to stdout there's no real reason to think anything would have
	// changed so return above. Otherwise, opportunistically try to re-open the
	// FD, once per call.
	b.f.Close()
	b.f = nil

	if err := b.open(); err != nil {
		b.fileLock.Unlock()
		return err
	}

	reader.Seek(0, io.SeekStart)
	_, err := reader.WriteTo(writer)
	b.fileLock.Unlock()
	return err
}

<<<<<<< HEAD
// Deprecated: Use eventlogger.
func (b *Backend) LogResponse(ctx context.Context, in *logical.LogInput) error {
	var writer io.Writer
	switch b.filePath {
=======
func (b *Backend) LogResponse(ctx context.Context, in *logical.LogInput) error {
	var writer io.Writer
	switch b.path {
>>>>>>> 4cb759cfc9 (fixed log)
	case stdout:
		writer = os.Stdout
	case discard:
		return nil
	}

	buf := bytes.NewBuffer(make([]byte, 0, 6000))
	err := b.formatter.FormatAndWriteResponse(ctx, buf, in)
	if err != nil {
		return err
	}

	return b.log(ctx, buf, writer)
}

func (b *Backend) LogTestMessage(ctx context.Context, in *logical.LogInput, config map[string]string) error {
	// Event logger behavior - manually Process each node
	if len(b.nodeIDList) > 0 {
		return audit.ProcessManual(ctx, in, b.nodeIDList, b.nodeMap)
	}

	// Old behavior
	var writer io.Writer
<<<<<<< HEAD
	switch b.filePath {
=======
	switch b.path {
>>>>>>> 4cb759cfc9 (fixed log)
	case stdout:
		writer = os.Stdout
	case discard:
		return nil
	}

	var buf bytes.Buffer

	temporaryFormatter, err := audit.NewTemporaryFormatter(config["format"], config["prefix"])
	if err != nil {
		return err
	}

	if err = temporaryFormatter.FormatAndWriteRequest(ctx, &buf, in); err != nil {
		return err
	}

	return b.log(ctx, &buf, writer)
}

// The file lock must be held before calling this
<<<<<<< HEAD
// Deprecated: Use eventlogger.
=======
>>>>>>> 4cb759cfc9 (fixed log)
func (b *Backend) open() error {
	if b.f != nil {
		return nil
	}
<<<<<<< HEAD
	if err := os.MkdirAll(filepath.Dir(b.filePath), b.mode); err != nil {
=======
	if err := os.MkdirAll(filepath.Dir(b.path), b.mode); err != nil {
>>>>>>> 4cb759cfc9 (fixed log)
		return err
	}

	var err error
<<<<<<< HEAD
	b.f, err = os.OpenFile(b.filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, b.mode)
=======
	b.f, err = os.OpenFile(b.path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, b.mode)
>>>>>>> 4cb759cfc9 (fixed log)
	if err != nil {
		return err
	}

	// Change the file mode in case the log file already existed. We special
	// case /dev/null since we can't chmod it and bypass if the mode is zero
<<<<<<< HEAD
	switch b.filePath {
	case "/dev/null":
	default:
		if b.mode != 0 {
			err = os.Chmod(b.filePath, b.mode)
=======
	switch b.path {
	case "/dev/null":
	default:
		if b.mode != 0 {
			err = os.Chmod(b.path, b.mode)
>>>>>>> 4cb759cfc9 (fixed log)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (b *Backend) Reload(_ context.Context) error {
	// When there are nodes created in the map, use the eventlogger behavior.
	if len(b.nodeMap) > 0 {
		for _, n := range b.nodeMap {
			if n.Type() == eventlogger.NodeTypeSink {
				return n.Reopen()
			}
		}

		return nil
	} else {
		// old non-eventlogger behavior
<<<<<<< HEAD
		switch b.filePath {
=======
		switch b.path {
>>>>>>> 4cb759cfc9 (fixed log)
		case stdout, discard:
			return nil
		}

		b.fileLock.Lock()
		defer b.fileLock.Unlock()

		if b.f == nil {
			return b.open()
		}

		err := b.f.Close()
		// Set to nil here so that even if we error out, on the next access open()
		// will be tried
		b.f = nil
		if err != nil {
			return err
		}

		return b.open()
	}
}

func (b *Backend) Invalidate(_ context.Context) {
	b.saltMutex.Lock()
	defer b.saltMutex.Unlock()
	b.salt.Store((*salt.Salt)(nil))
}

<<<<<<< HEAD
// formatterConfig creates the configuration required by a formatter node using
// the config map supplied to the factory.
func formatterConfig(config map[string]string) (audit.FormatterConfig, error) {
	const op = "file.formatterConfig"

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
	const op = "file.(Backend).configureFilterNode"

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
	const op = "file.(Backend).configureFormatterNode"

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
func (b *Backend) configureSinkNode(name string, filePath string, mode string, format string) error {
	const op = "file.(Backend).configureSinkNode"

	name = strings.TrimSpace(name)
	if name == "" {
		return fmt.Errorf("%s: name is required: %w", op, event.ErrInvalidParameter)
	}

	filePath = strings.TrimSpace(filePath)
	if filePath == "" {
		return fmt.Errorf("%s: file path is required: %w", op, event.ErrInvalidParameter)
	}

	format = strings.TrimSpace(format)
	if format == "" {
		return fmt.Errorf("%s: format is required: %w", op, event.ErrInvalidParameter)
	}

	sinkNodeID, err := event.GenerateNodeID()
	if err != nil {
		return fmt.Errorf("%s: error generating random NodeID for sink node: %w", op, err)
	}

	// normalize file path if configured for stdout or discard
	if strings.EqualFold(filePath, stdout) {
		filePath = stdout
	} else if strings.EqualFold(filePath, discard) {
		filePath = discard
	}

	var sinkNode eventlogger.Node
	var sinkName string

	switch filePath {
	case stdout:
		sinkName = stdout
		sinkNode, err = event.NewStdoutSinkNode(format)
	case discard:
		sinkName = discard
		sinkNode = event.NewNoopSink()
	default:
		// The NewFileSink function attempts to open the file and will return an error if it can't.
		sinkName = name
		sinkNode, err = event.NewFileSink(filePath, format, []event.Option{event.WithFileMode(mode)}...)
	}

	if err != nil {
		return fmt.Errorf("%s: file sink creation failed for path %q: %w", op, filePath, err)
	}

	sinkNode = &audit.SinkWrapper{Name: sinkName, Sink: sinkNode}

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
